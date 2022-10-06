package jenkins

import (
	"fmt"
	"github.com/argonsecurity/go-environments/logger"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/github"
	"github.com/argonsecurity/go-environments/environments/jenkins/environments"
	bitbucketserver "github.com/argonsecurity/go-environments/environments/jenkins/environments/bitbucket_server"
	"github.com/argonsecurity/go-environments/environments/jenkins/environments/gitlab"
	"github.com/argonsecurity/go-environments/environments/utils"
	"github.com/argonsecurity/go-environments/environments/utils/git"
	"github.com/argonsecurity/go-environments/http"
	"github.com/argonsecurity/go-environments/models"
)

const (
	builder           = "Jenkins"
	repositoryPathEnv = "WORKSPACE"

	jenkinsURLEnv         = "JENKINS_URL"
	buildURLEnv           = "BUILD_URL"
	runURLEnv             = "RUN_DISPLAY_URL"
	repositoryCloneURLEnv = "GIT_URL"

	buildIDEnv     = "BUILD_ID"
	buildNumberEnv = "BUILD_NUMBER"
	nodeIDEnv      = "NODE_NAME"

	jobNameEnv   = "JOB_NAME"
	nodeNameEnv  = "NODE_NAME"
	runNameEnv   = "BUILD_TAG"
	stageNameEnv = "STAGE_NAME"

	commitShaEnv     = "GIT_COMMIT"
	branchEnv        = "BRANCH_NAME"
	targetBranchName = "CHANGE_TARGET"

	githubHostname    = "github.com"
	gitlabHostname    = "gitlab.com"
	azureHostname     = "dev.azure.com"
	bitbucketHostname = "bitbucket.org"

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

// parseDataFromCloneUrl extracts data from the clone url
// and returns the repository url, organization and repository name
// the base url is used for cases where the base of the scm url includes a part of the URI
//
// i.e https://example.company.io/gitlab
func parseDataFromCloneUrl(cloneUrl, apiUrl string, repoSource enums.Source) (string, string, string, error) {
	var regexp = uriRegexp
	baseUrl, uri, isSshUrl, err := getUriFromCloneUrl(cloneUrl, apiUrl)
	if err != nil {
		return "", "", "", err
	}

	// In bitbucket server the clone url looks like this: https://server-bitbucket.company.com/scm/project/repo.git
	// so we need to extract the organization and repository names using a different regex
	if isSshUrl {
		regexp = sshUriRegexp
	} else if repoSource == enums.BitbucketServer {
		regexp = bitbucketServerUriRegexp
	}
	results := regexp.FindAllStringSubmatch(uri, -1)
	logger.Warnf("uri %s", uri)
	logger.Warnf("results %s", results)
	logger.Warnf("regexp %s", regexp)
	logger.Warnf("isSshUrl %b", isSshUrl)
	logger.Warnf("repoSource %b", repoSource)
	if len(results) == 0 {
		return "", "", "", fmt.Errorf("could not parse clone url: %s", cloneUrl)
	}
	result := results[0]

	var org, subgroups, repo string
	if len(result) == 4 { // url contains subgroups
		org, subgroups, repo = result[1], result[2], result[3]
	} else { // url doesn't contains subgroups
		org, repo = result[1], result[2]
	}

	return environments.BuildScmLink(baseUrl, org, subgroups, repo, isSshUrl, repoSource), org, repo, nil
}

// getUriFromCloneUrl for cases where the baseUrl is not actually
// a part of the cloneUrl (i.e. Github), we need to extract the URI
// from the cloneUrl without using the baseUrl
func getUriFromCloneUrl(cloneUrl, apiUrl string) (string, string, bool, error) {
	isSshUrl := sshIdentificationRegexp.MatchString(cloneUrl)
	if strings.Contains(cloneUrl, apiUrl) && apiUrl != "" {
		return apiUrl, strings.Replace(cloneUrl, apiUrl, "", 1), isSshUrl, nil
	}
	if httpUrlRegexp.MatchString(cloneUrl) {
		result := httpUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)
		return result[0][1], result[0][2], isSshUrl, nil
	}
	if sshUrlRegexp.MatchString(cloneUrl) {
		result := sshUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)
		return strings.Replace(result[0][1], "ssh", "https", 1), result[0][2], isSshUrl, nil
	}
	results := gitUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)
	if len(results) == 0 {
		return "", "", isSshUrl, fmt.Errorf("could not parse clone url: %s", cloneUrl)
	}
	result := results[0]
	return fmt.Sprintf("https://%s", result[1]), strings.Replace(result[2], ":", "/", 1), isSshUrl, nil
}

func loadConfiguration() (*models.Configuration, error) {
	repositoryPath := os.Getenv(repositoryPathEnv)
	if !git.IsPathContainsRepository(repositoryPath) {
		return nil, fmt.Errorf("%s is not a repository path", repositoryPath)
	}

	cloneUrl, err := getRepositoryCloneURL(repositoryPath)
	cloneUrl = utils.StripCredentialsFromUrl(cloneUrl)
	if err != nil {
		return nil, err
	}

	repoSource, apiUrl := getRepositorySource(cloneUrl)
	repositoryURL, org, repositoryName, err := parseDataFromCloneUrl(cloneUrl, apiUrl, repoSource)
	if err != nil {
		return nil, err
	}

	commit := os.Getenv(commitShaEnv)
	if commit == "" {
		commit, err = git.GetGitCommit(repositoryPath)
		if err != nil {
			return nil, err
		}
	}

	scmId := utils.GenerateScmId(cloneUrl)

	branch := getBranchName(repositoryPath, commit)
	configuration := &models.Configuration{
		Url:       os.Getenv(jenkinsURLEnv),
		SCMApiUrl: apiUrl,
		LocalPath: repositoryPath,
		Branch:    branch,
		CommitSha: commit,
		Repository: models.Repository{
			Name:     repositoryName,
			CloneUrl: cloneUrl,
			Source:   repoSource,
			Url:      repositoryURL,
		},
		Pipeline: models.Entity{
			Id:   os.Getenv(jobNameEnv),
			Name: os.Getenv(jobNameEnv),
		},
		Job: models.Entity{
			Id:   os.Getenv(stageNameEnv),
			Name: os.Getenv(stageNameEnv),
		},
		Run: models.BuildRun{
			BuildId:     os.Getenv(buildIDEnv),
			BuildNumber: os.Getenv(buildNumberEnv),
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
			TargetRef: models.Ref{
				Branch: os.Getenv(targetBranchName),
			},
		},
		Builder: builder,
		Organization: models.Entity{
			Name: org,
		},
		PipelinePaths: getAllPipelinePaths(repositoryPath),
		Environment:   enums.Jenkins,
		ScmId:         scmId,
	}

	configuration = environments.EnhanceConfiguration(configuration)
	if configuration.Pusher.Username == "" {
		configuration.Pusher.Username = utils.DetectPusher()
	}
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
	return os.Getenv(runURLEnv)
}

func (env environment) GetBuildLink() string {
	url := os.Getenv(buildURLEnv)
	if url != "" {
		return url
	}
	return os.Getenv(runURLEnv)
}

func (e environment) GetFileLineLink(filename string, ref string, line int) string {
	return ""
}

func (e environment) Name() string {
	return "jenkins"
}

func (e environment) IsCurrentEnvironment() bool {
	var isExist bool
	if _, isExist = os.LookupEnv("JENKINS_HOME"); !isExist {
		_, isExist = os.LookupEnv("JENKINS_URL")
	}
	return isExist
}

func getRepositoryCloneURL(repositoryPath string) (string, error) {
	if cloneUrl, isExist := os.LookupEnv(repositoryCloneURLEnv); isExist {
		return cloneUrl, nil
	}
	return git.GetGitRemoteURL(repositoryPath)
}

func getRepositorySource(cloneUrl string) (enums.Source, string) {
	switch {
	case strings.Contains(cloneUrl, bitbucketHostname):
		return enums.Bitbucket, bitbucketApiUrl
	case strings.Contains(cloneUrl, githubHostname):
		return enums.Github, githubApiUrl
	case strings.Contains(cloneUrl, azureHostname):
		return enums.Azure, azureApiUrl
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
