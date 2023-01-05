package circleci

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/argonsecurity/go-environments/environments/circleci/environments/github"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/utils"
	"github.com/argonsecurity/go-environments/models"
)

const (
	builder               = "CircleCi"
	buildNumberEnv        = "CIRCLE_BUILD_NUM"
	repositoryCloneURLEnv = "CIRCLE_REPOSITORY_URL"
	commitShaEnv          = "CIRCLE_SHA1"
	repositoryNameEnv     = "CIRCLE_PROJECT_REPONAME"
	branchEnv             = "CIRCLE_BRANCH"
	circlePullRequestUrl  = "CIRCLE_PULL_REQUEST"
	workflowIdEnv         = "CIRCLE_WORKFLOW_ID"
	jobNameEnv            = "CIRCLE_JOB"
	jobIdEnv              = "CIRCLE_WORKFLOW_JOB_ID"
	buildUrlEnv           = "CIRCLE_BUILD_URL"
	pipelinePath          = ".circleci/config.yml"

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
	CircleCi      = environment{}
	configuration *models.Configuration
)

type environment struct{}

func (e environment) GetBuildLink() string {
	return os.Getenv(buildUrlEnv)
}

func (e environment) GetStepLink() string {
	return ""
}

func (e environment) GetFileLink(filename string, ref string, commit string) string {
	return ""
}

func (e environment) GetFileLineLink(filename string, ref string, commit string, startLine int, endLine int) string {
	return ""
}

func (e environment) GetConfiguration() (*models.Configuration, error) {
	if configuration == nil {
		loadedConfiguration, err := loadConfiguration()
		configuration = loadedConfiguration
		return configuration, err
	}
	return configuration, nil
}

func loadConfiguration() (*models.Configuration, error) {
	repoCloneUrl := os.Getenv(repositoryCloneURLEnv)
	source, apiUrl := GetRepositorySource(repoCloneUrl)
	scmUrl, org, _, repositoryFullName, err := utils.ParseDataFromCloneUrl(repoCloneUrl, apiUrl, source)
	scmLink := scmUrl
	if !strings.HasSuffix(scmLink, ".git") {
		scmLink += ".git"
	}
	if err != nil {
		return nil, err
	}
	scmId := utils.GenerateScmId(scmLink)

	pullRequestUrl := os.Getenv(circlePullRequestUrl)
	pullRequestId := path.Base(pullRequestUrl)
	prNumber, _ := strconv.Atoi(pullRequestId)
	if err != nil {
		return nil, err
	}
	targetBranch, _ := github.GetPullRequestTargetBranch(org, os.Getenv(repositoryNameEnv), prNumber)

	configuration = &models.Configuration{
		Url:       "https://app.circleci.com",
		SCMApiUrl: apiUrl,
		LocalPath: repoCloneUrl,
		Branch:    os.Getenv(branchEnv),
		CommitSha: os.Getenv(commitShaEnv),
		Repository: models.Repository{
			Name:     os.Getenv(repositoryNameEnv),
			FullName: repositoryFullName,
			CloneUrl: scmLink,
			Source:   source,
			Url:      scmUrl,
		},
		Pipeline: models.Pipeline{
			Entity: models.Entity{
				Id:   os.Getenv(workflowIdEnv),
				Name: os.Getenv(workflowIdEnv),
			},
			Path: getPipelinePath(),
		},
		Job: models.Entity{
			Id:   os.Getenv(jobIdEnv),
			Name: os.Getenv(jobNameEnv),
		},
		Builder: builder,
		Organization: models.Entity{
			Name: org,
		},
		Run: models.BuildRun{
			BuildId:     os.Getenv(buildNumberEnv),
			BuildNumber: os.Getenv(buildNumberEnv),
		},
		Pusher: models.Pusher{
			Username: utils.DetectPusher(),
		},
		Runner: models.Runner{
			OS:           runtime.GOOS,
			Architecture: runtime.GOARCH,
		},
		PullRequest: models.PullRequest{
			Id: pullRequestId,
			SourceRef: models.Ref{
				Branch: os.Getenv(branchEnv),
			},
			TargetRef: models.Ref{
				Branch: targetBranch,
			},
		},
		Environment:   enums.CircleCi,
		ScmId:         scmId,
		PipelinePaths: []string{getPipelinePath()},
	}

	return configuration, nil
}

func getPipelinePath() string {
	if _, err := os.Stat(pipelinePath); err == nil {
		return pipelinePath
	}
	return ""
}

func (e environment) Name() string {
	return "circleci"
}

func (e environment) IsCurrentEnvironment() bool {
	circleCi := os.Getenv("CIRCLECI")
	return circleCi == "true"
}

func GetRepositorySource(cloneUrl string) (enums.Source, string) {
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

	return enums.Unknown, ""
}
