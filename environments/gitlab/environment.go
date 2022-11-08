package gitlab

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/utils"
	"github.com/argonsecurity/go-environments/models"
)

const (
	jobIdEnv   = "CI_JOB_ID"
	jobNameEnv = "CI_JOB_NAME"

	repositoryPathEnv     = "CI_PROJECT_DIR"
	projectNameEnv        = "CI_PROJECT_NAME"
	groupNameEnv          = "CI_PROJECT_NAMESPACE"
	projectIdEnv          = "CI_PROJECT_ID"
	projectUrlEnv         = "CI_PROJECT_URL"
	rootNamespaceEnv      = "CI_PROJECT_ROOT_NAMESPACE"
	repositoryCloneURLEnv = "CI_REPOSITORY_URL"
	commitAuthorEnv       = "CI_COMMIT_AUTHOR"

	runnerIdEnv          = "CI_RUNNER_ID"
	runnerOSEnv          = "CI_RUNNER_EXECUTABLE_ARCH"
	runnerDescriptionEnv = "CI_RUNNER_DESCRIPTION"

	pipelineFilePathEnv = "CI_CONFIG_PATH"

	commitShaEnv       = "CI_COMMIT_SHA"
	beforeCommitShaEnv = "CI_COMMIT_BEFORE_SHA"
	branchEnv          = "CI_COMMIT_REF_NAME"

	mergeRequestIdEnv     = "CI_MERGE_REQUEST_ID"
	mergeSourceBranchSha  = "CI_MERGE_REQUEST_SOURCE_BRANCH_SHA"
	mergeSourceBranchName = "CI_MERGE_REQUEST_SOURCE_BRANCH_NAME"
	mergeTargetBranchSha  = "CI_MERGE_REQUEST_TARGET_BRANCH_SHA"
	mergeTargetBranchName = "CI_MERGE_REQUEST_TARGET_BRANCH_NAME"

	pipelineIdEnv = "CI_PIPELINE_ID"
	gitlabUrlEnv  = "CI_SERVER_URL"
)

var (
	Gitlab        = environment{}
	configuration *models.Configuration

	gitlabPipelines = []string{".gitlab-ci.yml", ".gitlab-ci.yaml"}
)

type environment struct{}

func (e environment) GetConfiguration() (*models.Configuration, error) {
	if configuration == nil {
		loadConfiguration()
	}
	return configuration, nil
}

func loadConfiguration() *models.Configuration {
	source := getSource()
	repoPath := os.Getenv(repositoryPathEnv)
	cloneUrl := utils.StripCredentialsFromUrl(os.Getenv(repositoryCloneURLEnv))
	scmId := utils.GenerateScmId(cloneUrl)

	configuration = &models.Configuration{
		Url:             os.Getenv(gitlabUrlEnv),
		SCMApiUrl:       os.Getenv(gitlabUrlEnv),
		LocalPath:       repoPath,
		Branch:          os.Getenv(branchEnv),
		CommitSha:       os.Getenv(commitShaEnv),
		BeforeCommitSha: os.Getenv(beforeCommitShaEnv),
		Organization: models.Entity{
			Name: os.Getenv(rootNamespaceEnv),
		},
		Repository: models.Repository{
			Id:       os.Getenv(projectIdEnv),
			Name:     os.Getenv(projectNameEnv),
			Url:      os.Getenv(projectUrlEnv),
			CloneUrl: cloneUrl,
			Source:   source,
		},
		Pipeline: models.Pipeline{
			Entity: models.Entity{
				Id:   os.Getenv(pipelineIdEnv), // Each GitLab project has only one pipeline
				Name: os.Getenv(projectNameEnv),
			},
			Path: os.Getenv(pipelineFilePathEnv),
		},
		Job: models.Entity{
			Id:   os.Getenv(jobNameEnv),
			Name: os.Getenv(jobNameEnv),
		},
		Run: models.BuildRun{
			BuildId: os.Getenv(jobIdEnv),
		},
		Runner: models.Runner{
			Id:           os.Getenv(runnerIdEnv),
			Name:         os.Getenv(runnerDescriptionEnv),
			OS:           os.Getenv(runnerOSEnv),
			Architecture: runtime.GOARCH,
		},
		PullRequest: models.PullRequest{
			Id: os.Getenv(mergeRequestIdEnv),
			SourceRef: models.Ref{
				Branch: os.Getenv(mergeSourceBranchName),
				Sha:    os.Getenv(mergeSourceBranchSha),
			},
			TargetRef: models.Ref{
				Branch: os.Getenv(mergeTargetBranchName),
				Sha:    os.Getenv(mergeTargetBranchSha),
			},
		},
		Pusher: models.Pusher{
			Username: getUsername(),
		},
		PipelinePaths: getPipelinePaths(repoPath),
		Environment:   source,
		ScmId:         scmId,
	}
	return configuration
}

func (e environment) GetStepLink() string {
	return fmt.Sprintf("%s/%s/%s/-/jobs/%s", os.Getenv(gitlabUrlEnv), os.Getenv(groupNameEnv), os.Getenv(projectNameEnv), os.Getenv(jobIdEnv))
}

func (e environment) GetBuildLink() string {
	return fmt.Sprintf("%s/%s/%s/-/pipelines/%s", os.Getenv(gitlabUrlEnv), os.Getenv(groupNameEnv), os.Getenv(projectNameEnv), os.Getenv(pipelineIdEnv))
}

func (e environment) GetFileLink(filename string, branch string, commit string) string {
	repoURL := os.Getenv(projectUrlEnv)
	return GetFileLink(
		repoURL,
		filename,
		branch,
		commit,
	)
}

func (e environment) GetFileLineLink(filename string, branch string, commit string, startLine int, endLine int) string {
	repoURL := os.Getenv(projectUrlEnv)
	return GetFileLineLink(
		repoURL,
		filename,
		branch,
		commit,
		startLine,
		endLine,
	)
}

func GetFileLink(repositoryURL string, filename, branch string, commit string) string {
	refToUse := branch
	if commit != "" {
		refToUse = commit
	}

	return fmt.Sprintf("%s/-/blob/%s/%s",
		repositoryURL,
		refToUse,
		filename,
	)
}

func GetFileLineLink(repositoryURL string, filename string, branch string, commit string, startLine, endLine int) string {
	url := GetFileLink(repositoryURL, filename, branch, commit)
	if startLine != 0 {
		if endLine == 0 {
			endLine = startLine
		}

		url = fmt.Sprintf("%s#L%d-%d", url, startLine, endLine)
	}
	return url
}

func (e environment) Name() string {
	return "gitlab"
}

func getSource() enums.Source {
	url := os.Getenv(gitlabUrlEnv)
	if url == "https://gitlab.com" {
		return enums.Gitlab
	}
	return enums.GitlabServer
}

func (e environment) IsCurrentEnvironment() bool {
	_, isExist := os.LookupEnv("GITLAB_CI")
	return isExist
}

func getPipelinePaths(rootDir string) []string {
	paths := make([]string, 0)

	for _, pipelineFile := range gitlabPipelines {
		path := filepath.Join(rootDir, pipelineFile)
		if _, err := os.Stat(path); err == nil {
			paths = append(paths, path)
		}
	}

	return paths
}

func getUsername() string {
	arr := strings.Split(os.Getenv(commitAuthorEnv), " ")
	username := strings.Join(arr[:len(arr)-1][:], " ")
	if username == "" {
		username = utils.DetectPusher()
	}
	return username
}
