package jenkins

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/argonsecurity/go-utils/environments/enums"
	"github.com/argonsecurity/go-utils/environments/environments/github"
	"github.com/argonsecurity/go-utils/environments/environments/jenkins/environments"
	bitbucketserver "github.com/argonsecurity/go-utils/environments/environments/jenkins/environments/bitbucket_server"
	"github.com/argonsecurity/go-utils/environments/environments/jenkins/environments/gitlab"
	"github.com/argonsecurity/go-utils/environments/environments/utils"
	"github.com/argonsecurity/go-utils/environments/environments/utils/git"
	"github.com/argonsecurity/go-utils/environments/models"
	"github.com/argonsecurity/go-utils/http"
)

const (
	builder           = "circleCI"
	repositoryPathEnv = "CIRCLE_WORKING_DIRECTORY"

	buildURLEnv           = "CIRCLE_BUILD_URL"
	repositoryCloneURLEnv = "CIRCLE_REPOSITORY_URL"
	repositoryNameEnv     = "CIRCLE_PROJECT_REPONAME"

	nodeIDEnv = "CIRCLE_NODE_INDEX"

	jobNameEnv     = "CIRCLE_JOB"
	nodeNameEnv    = "CIRCLE_NODE_INDEX"
	buildNumberEnv = "CIRCLE_BUILD_NUM"
	buildIdEnv     = "CIRCLE_BUILD_NUM"
	authourNameEnv = "CIRCLE_PR_USERNAME"
	stageNameEnv   = "CIRCLE_WORKFLOW_JOB_ID"
	ownerNameEnv   = "CIRCLE_PROJECT_USERNAME"

	commitShaEnv     = "CIRCLE_SHA1"
	branchEnv        = "CIRCLE_BRANCH"
	targetBranchName = "CHANGE_TARGET"

	githubHostname    = "github.com"
	gitlabHostname    = "gitlab.com"
	azureHostname     = "dev.azure.com"
	bitbucketHostname = "	.org"

	githubApiUrl    = "https://api.github.com"
	gitlabApiUrl    = "https://gitlab.com/api/v4"
	azureApiUrl     = ""
	bitbucketApiUrl = "https://api.bitbucket.org/2.0"
)

var (
	Jenkins       = environment{}
	configuration *models.Configuration

	bitbucketServerUriRegexp = regexp.MustCompile(`scm/(.*?)/(.*?)(?:\.git|$)`)
	uriRegexp                = regexp.MustCompile(`/?(.+?)/(?:(.+/))?(.+?)(?:\.git|$)`)
	httpUrlRegexp            = regexp.MustCompile(`(https?://.+?)(/.+)`)
	sshUrlRegexp             = regexp.MustCompile(`(ssh?://.+?)(?:\:[0-9]+)(/.+)`)
	gitUrlRegexp             = regexp.MustCompile(`git@(.+?)(\:.+)`)
	sshUriRegexp             = regexp.MustCompile(`(?:/(?:v3|[0-9]+))?/(?P<org>.+?)/(.+/)?(?P<repo>.+?)(?:\.git|$)`)
	sshIdentificationRegexp  = regexp.MustCompile(`^.*@|ssh://`)
)

type environment struct{}

func (e environment) GetConfiguration() (*models.Configuration, error) {
	if configuration == nil {
		loadedConfiguration, err := loadConfiguration()
		configuration = loadedConfiguration
		return configuration, err
	}
	return configuration, nil
}

func getUriFromCloneUrl(cloneUrl, apiUrl string) (string, string, bool) {
	isSshUrl := sshIdentificationRegexp.MatchString(cloneUrl)
	if strings.Contains(cloneUrl, apiUrl) && apiUrl != "" {
		return apiUrl, strings.Replace(cloneUrl, apiUrl, "", 1), isSshUrl
	}
	if httpUrlRegexp.MatchString(cloneUrl) {
		result := httpUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)
		return result[0][1], result[0][2], isSshUrl
	}
	if sshUrlRegexp.MatchString(cloneUrl) {
		result := sshUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)
		return strings.Replace(result[0][1], "ssh", "https", 1), result[0][2], isSshUrl
	}
	result := gitUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)[0]
	return fmt.Sprintf("https://%s", result[1]), strings.Replace(result[2], ":", "/", 1), isSshUrl
}

func parseDataFromCloneUrl(cloneUrl, apiUrl, repo, org string, repoSource enums.Source) string {
	var regexp = uriRegexp
	baseUrl, uri, isSshUrl := getUriFromCloneUrl(cloneUrl, apiUrl)

	// In bitbucket server the clone url looks like this: https://server-bitbucket.company.com/scm/project/repo.git
	// so we need to extract the organization and repository names using a different regex
	if isSshUrl {
		regexp = sshUriRegexp
	} else if repoSource == enums.BitbucketServer {
		regexp = bitbucketServerUriRegexp
	}
	result := regexp.FindAllStringSubmatch(uri, -1)[0]

	var subgroups string
	if len(result) == 4 { // url contains subgroups
		subgroups = result[2]
	}

	return environments.BuildScmLink(baseUrl, org, subgroups, repo, isSshUrl, repoSource)
}

func parseDataFromBuildUrl(buildUrl string, source enums.Source) string {
	dictionary := map[enums.Source]interface{}{
		enums.Github:    "gh",
		enums.Gitlab:    "gl",
		enums.Bitbucket: "bb",
	}

	return strings.Split(buildUrl, fmt.Sprintf("/%s", dictionary[source]))[0]
}

func loadConfiguration() (*models.Configuration, error) {
	cloneUrl := os.Getenv(repositoryCloneURLEnv)
	cloneUrl = utils.StripCredentialsFromUrl(cloneUrl)

	repoSource, apiUrl := getRepositorySource(cloneUrl)
	repositoryName := os.Getenv(repositoryNameEnv)
	org := os.Getenv(ownerNameEnv)
	repositoryURL := parseDataFromCloneUrl(cloneUrl, apiUrl, repositoryName, org, repoSource)
	buildUrl := os.Getenv(buildURLEnv)
	circleCiURL := parseDataFromBuildUrl(buildUrl, repoSource)

	commit := os.Getenv(commitShaEnv)

	scmId := utils.GenerateScmId(cloneUrl)
	branch := os.Getenv(branchEnv)

	configuration := &models.Configuration{
		Url:       circleCiURL,
		SCMApiUrl: apiUrl,
		LocalPath: os.Getenv(repositoryPathEnv),
		Branch:    branch,
		CommitSha: commit,
		Repository: models.Repository{
			Name:     repositoryName,
			CloneUrl: cloneUrl,
			Source:   repoSource,
			Url:      repositoryURL,
		},
		// Pipeline: models.Entity{
		// 	Id:   os.Getenv(jobNameEnv),
		// 	Name: os.Getenv(jobNameEnv),
		// },
		// Job: models.Entity{
		// 	Id:   os.Getenv(stageNameEnv),
		// 	Name: os.Getenv(stageNameEnv),
		// },
		Run: models.Entity{
			Id:   os.Getenv(buildNumberEnv),
			Name: os.Getenv(buildNumberEnv),
		},
		Runner: models.Runner{
			Id:           os.Getenv(nodeIDEnv),
			Name:         os.Getenv(nodeNameEnv),
			OS:           runtime.GOOS,
			Architecture: runtime.GOARCH,
		},
		PullRequest: models.PullRequest{
			SourceRef: models.Ref{
				Branch: branch,
			},
			// TargetRef: models.Ref{
			// 	Branch: branch,
			// },
		},
		Builder: builder,
		Organization: models.Entity{
			Name: org,
		},
		//PipelinePaths: getAllPipelinePaths(""),
		Environment: enums.CircleCi,
		ScmId:       scmId,
	}

	configuration = environments.EnhanceConfiguration(configuration)
	configuration.Repository.CloneUrl = utils.StripCredentialsFromUrl(configuration.Repository.CloneUrl)
	return configuration, nil
}

func getBranchName(repositoryPath string, commit string) string {
	branchName := os.Getenv(branchEnv)
	if branchName == "" {
		branchName, _ = git.GetGitBranch(repositoryPath, commit)
	}
	return branchName
}

func (env environment) GetStepLink() string {
	return ""
}

func (env environment) GetBuildLink() string {
	return os.Getenv(buildURLEnv)
}

func (e environment) GetFileLineLink(filename string, ref string, line int) string {
	return ""
}

func (e environment) Name() string {
	return builder
}

func getRepositorySource(cloneUrl string) (enums.Source, string) {
	switch {
	case strings.Contains(cloneUrl, bitbucketHostname):
		return enums.Bitbucket, bitbucketApiUrl
	case strings.Contains(cloneUrl, githubHostname):
		return enums.Github, githubApiUrl
	case strings.Contains(cloneUrl, gitlabHostname):
		return enums.Gitlab, gitlabApiUrl

	}

	return discoverSCMSource(cloneUrl)
}

func discoverSCMSource(gitUrl string) (enums.Source, string) {
	urls := utils.ParseGitURL(gitUrl)
	httpClient := http.GetHTTPClient("", nil)
	for _, url := range urls {
		if gitlab.CheckGitlabByHTTPRequest(url, httpClient) {
			return enums.GitlabServer, url
		}

		if bitbucketserver.CheckBitbucketServerByHTTPRequest(url, httpClient) {
			return enums.BitbucketServer, url
		}
	}
	return enums.Unknown, ""
}

func getJenkinsPipelinePaths(rootDir string) string {
	jenkinsfilePath := filepath.Join(rootDir, "Jenkinsfile")
	if _, err := os.Stat(jenkinsfilePath); !os.IsNotExist(err) {
		return jenkinsfilePath
	}
	return ""
}

func getAllPipelinePaths(rootDir string) []string {
	paths := make([]string, 0)
	if jenkinsfilePath := getJenkinsPipelinePaths(rootDir); jenkinsfilePath != "" {
		paths = append(paths, jenkinsfilePath)
	}
	paths = append(paths, github.GetPipelinePaths(rootDir)...)

	return paths
}
