package bitbucket

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/utils"
	"github.com/argonsecurity/go-environments/models"
)

const (
	repositoryPathEnv      = "BITBUCKET_CLONE_DIR"
	repositoryNameEnv      = "BITBUCKET_REPO_SLUG"
	repositoryIdEnv        = "BITBUCKET_REPO_UUID"
	repositoryUrlEnv       = "BITBUCKET_GIT_HTTP_ORIGIN"
	repositoryFullNameEnv  = "BITBUCKET_REPO_FULL_NAME"
	workspaceEnv           = "BITBUCKET_WORKSPACE"
	prDestentaionBranchEnv = "BITBUCKET_PR_DESTINATION_BRANCH"

	buildNumber = "BITBUCKET_BUILD_NUMBER"

	commitShaEnv = "BITBUCKET_COMMIT"
	branchEnv    = "BITBUCKET_BRANCH"

	mergeRequestIdEnv = "BITBUCKET_PR_ID"
	pipelineIdEnv     = "BITBUCKET_PIPELINE_UUID"
	stepIdEnv         = "BITBUCKET_STEP_UUID"
)

var (
	bitbucketApiUrl = "https://api.bitbucket.org/2.0"
	bitbucketUrl    = "https://bitbucket.org"
)

var (
	Bitbucket     = environment{}
	configuration *models.Configuration

	bitbucketPipelines = []string{"bitbucket-pipelines.yml"}
)

type environment struct{}

func (e environment) GetConfiguration() (*models.Configuration, error) {
	if configuration == nil {
		loadConfiguration()
	}
	return configuration, nil
}

func loadConfiguration() *models.Configuration {
	source := enums.Bitbucket
	repoPath := os.Getenv(repositoryPathEnv)
	cloneUrl := fmt.Sprintf("%s.git", os.Getenv(repositoryUrlEnv))
	strippedCloneUrl := utils.StripCredentialsFromUrl(cloneUrl)
	scmId := utils.GenerateScmId(strippedCloneUrl)

	configuration = &models.Configuration{
		Url:       bitbucketUrl,
		SCMApiUrl: bitbucketApiUrl,
		LocalPath: repoPath,
		Branch:    os.Getenv(branchEnv),
		CommitSha: os.Getenv(commitShaEnv),
		Repository: models.Repository{
			Id:       os.Getenv(repositoryIdEnv),
			Name:     os.Getenv(repositoryNameEnv),
			Url:      os.Getenv(repositoryUrlEnv),
			CloneUrl: strippedCloneUrl,
			Source:   source,
		},
		Organization: models.Entity{
			Name: os.Getenv(workspaceEnv),
		},
		Pipeline: models.Entity{
			Id:   os.Getenv(pipelineIdEnv),
			Name: os.Getenv(repositoryNameEnv),
		},
		Run: models.BuildRun{
			BuildId:     os.Getenv(buildNumber),
			BuildNumber: os.Getenv(buildNumber),
		},
		Pusher: models.Pusher{
			Username: utils.DetectPusher(),
		},
		Runner: models.Runner{
			OS:           runtime.GOOS,
			Architecture: runtime.GOARCH,
		},
		PullRequest: models.PullRequest{
			Id: os.Getenv(mergeRequestIdEnv),
			TargetRef: models.Ref{
				Branch: os.Getenv(prDestentaionBranchEnv),
			},
		},
		PipelinePaths: getPipelinePaths(repoPath),
		Environment:   source,
		ScmId:         scmId,
	}
	return configuration
}

func (e environment) GetStepLink() string {
	return fmt.Sprintf("%s/%s/pipelines/results/%s/steps/%s", bitbucketUrl, os.Getenv(repositoryFullNameEnv), os.Getenv(buildNumber), os.Getenv(stepIdEnv))
}

func (e environment) GetBuildLink() string {
	return fmt.Sprintf("%s/%s/pipelines/results/%s", bitbucketUrl, os.Getenv(repositoryFullNameEnv), url.PathEscape(os.Getenv(buildNumber)))
}

func (e environment) GetFileLineLink(filename string, branch string, commit string, startLine int, endLine int) string {
	return GetFileLink(
		fmt.Sprintf("%s/%s", bitbucketUrl, os.Getenv(repositoryFullNameEnv)),
		filename,
		branch,
		commit,
		startLine,
		endLine,
	)
}

func GetFileLink(repositoryURL string, filename string, branch string, commit string, startLine, endLine int) string {
	link := ""
	if branch != "" && commit != "" {
		link = fmt.Sprintf("%s/src/%s/%s?at=%s",
			repositoryURL,
			commit,
			filename,
			url.PathEscape(branch),
		)
	} else if branch != "" && commit == "" {
		link = fmt.Sprintf("%s/src/%s/%s",
			repositoryURL,
			branch,
			filename,
		)
	} else {
		link = fmt.Sprintf("%s/src/%s/%s",
			repositoryURL,
			commit,
			filename,
		)
	}

	if startLine != 0 {
		if endLine == 0 {
			endLine = startLine
		}
		return fmt.Sprintf("%s#lines-%d:%d", link, startLine, endLine)
	}
	return link
}

func (e environment) IsCurrentEnvironment() bool {
	_, isExist := os.LookupEnv("BITBUCKET_PROJECT_KEY")
	return isExist
}

func (e environment) Name() string {
	return "bitbucket"
}

func getPipelinePaths(rootDir string) []string {
	paths := make([]string, 0)

	for _, pipelineFile := range bitbucketPipelines {
		path := filepath.Join(rootDir, pipelineFile)
		if _, err := os.Stat(path); err == nil {
			paths = append(paths, path)
		}
	}

	return paths
}
