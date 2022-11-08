package github

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/utils"
	"github.com/argonsecurity/go-environments/environments/utils/git"
	"github.com/argonsecurity/go-environments/models"
)

const (
	builder             = "Github Action"
	githubRepositoryEnv = "GITHUB_REPOSITORY"
	githubServerEnv     = "GITHUB_SERVER_URL"
	githubWorkflowEnv   = "GITHUB_WORKFLOW"
	githubRunIdEnv      = "GITHUB_RUN_ID"
	githubRunNumberEnv  = "GITHUB_RUN_NUMBER"
	repositoryPathEnv   = "GITHUB_WORKSPACE"

	githubJobEnv = "GITHUB_JOB"

	branchEnv    = "GITHUB_REF"
	commitShaEnv = "GITHUB_SHA"

	runnerNameEnv = "RUNNER_NAME"
	runnerOSEnv   = "RUNNER_OS"

	baseBranchNameEnv = "GITHUB_BASE_REF"
	headBranchNameEnv = "GITHUB_HEAD_REF"

	githubEventPath    = "GITHUB_EVENT_PATH"
	githubEventNameEnv = "GITHUB_EVENT_NAME"

	githubDir    = ".github"
	workflowsDir = ".github/workflows"

	githubApiUrlEnv = "GITHUB_API_URL"

	pullRequestEventName = "pull_request"
)

var (
	// Github environment
	Github        = environment{}
	configuration *models.Configuration
)

type environment struct{}

func (e environment) GetConfiguration() (*models.Configuration, error) {
	if configuration == nil {
		if err := loadConfiguration(); err != nil {
			return nil, err
		}
	}
	return configuration, nil
}

func loadConfiguration() error {
	payload, err := initPayload()
	if err != nil {
		return err
	}

	source := getSource()

	repoPath := os.Getenv(repositoryPathEnv)
	repoUrl := fmt.Sprintf("%s/%s", os.Getenv(githubServerEnv), os.Getenv(githubRepositoryEnv))
	cloneUrl, err := git.GetGitRemoteURL(repoPath)

	if err != nil || cloneUrl == "" {
		cloneUrl = fmt.Sprintf("%s.git", repoUrl)
	}

	username := payload.Sender.Login
	if username == "" {
		configuration.Pusher.Username = utils.DetectPusher()
	}
	strippedCloneUrl := utils.StripCredentialsFromUrl(cloneUrl)
	scmId := utils.GenerateScmId(strippedCloneUrl)

	pipelines := GetPipelinePaths(repoPath)
	repoId := strconv.Itoa(payload.Repository.Id)
	configuration = &models.Configuration{
		Url:       os.Getenv(githubServerEnv),
		SCMApiUrl: os.Getenv(githubApiUrlEnv),
		LocalPath: repoPath,
		CommitSha: os.Getenv(commitShaEnv),
		Branch:    getBranch(),
		Run: models.BuildRun{
			BuildId:     os.Getenv(githubRunIdEnv),
			BuildNumber: os.Getenv(githubRunNumberEnv),
		},
		Job: models.Entity{
			Id:   os.Getenv(githubJobEnv),
			Name: os.Getenv(githubJobEnv),
		},
		Pipeline: models.Pipeline{
			Entity: models.Entity{
				Id:   os.Getenv(githubWorkflowEnv),
				Name: os.Getenv(githubWorkflowEnv),
			},
			Path: getPipelinePath(os.Getenv(githubWorkflowEnv)),
		},
		Runner: models.Runner{
			Id:           os.Getenv(githubRunIdEnv),
			Name:         os.Getenv(runnerNameEnv),
			OS:           os.Getenv(runnerOSEnv),
			Architecture: runtime.GOARCH,
		},
		Repository: models.Repository{
			Id:       repoId,
			Name:     strings.Split(os.Getenv(githubRepositoryEnv), "/")[1],
			Url:      repoUrl,
			CloneUrl: strippedCloneUrl,
			Source:   source,
		},
		PullRequest: models.PullRequest{
			SourceRef: models.Ref{
				Branch: os.Getenv(headBranchNameEnv),
			},
			TargetRef: models.Ref{
				Branch: os.Getenv(baseBranchNameEnv),
			},
		},
		Commits: GetCommits(payload),
		Builder: builder,
		Organization: models.Entity{
			Name: payload.Repository.Owner.Login,
		},
		Pusher: models.Pusher{
			Entity: models.Entity{
				Id:   strconv.Itoa(payload.Sender.Id),
				Name: payload.Sender.Login,
			},
			Username: username,
		},
		PipelinePaths: pipelines,
		Environment:   source,
		ScmId:         scmId,
	}

	return nil
}

func getPipelinePath(githubWorkflow string) string {
	if strings.HasPrefix(githubWorkflow, ".github/workflows/") {
		return githubWorkflow
	}
	return ""
}

func getSource() enums.Source {
	url := os.Getenv(githubServerEnv)
	if url == "https://github.com" {
		return enums.Github
	}
	return enums.GithubServer
}

func getBranch() string {
	if os.Getenv(githubEventNameEnv) == pullRequestEventName {
		return os.Getenv(headBranchNameEnv)
	}
	return os.Getenv(branchEnv)
}

func (e environment) GetStepLink() string {
	return fmt.Sprintf("%s/%s/actions/runs/%s", os.Getenv(githubServerEnv), os.Getenv(githubRepositoryEnv), os.Getenv(githubRunIdEnv))
}

func (e environment) GetBuildLink() string {
	return fmt.Sprintf("%s/%s/actions/runs/%s", os.Getenv(githubServerEnv), os.Getenv(githubRepositoryEnv), os.Getenv(githubRunIdEnv))
}

func (e environment) GetFileLink(filename string, branch string, commit string) string {
	return GetFileLink(
		fmt.Sprintf("%s/%s", os.Getenv(githubServerEnv), os.Getenv(githubRepositoryEnv)),
		filename,
		branch,
		commit,
	)
}

func (e environment) GetFileLineLink(filename string, branch string, commit string, startLine int, endLine int) string {
	return GetFileLineLink(
		fmt.Sprintf("%s/%s", os.Getenv(githubServerEnv), os.Getenv(githubRepositoryEnv)),
		filename,
		branch,
		commit,
		startLine,
		endLine,
	)
}

func GetFileLink(repositoryURL string, filename string, branch string, commit string) string {
	refToUse := branch
	if commit != "" {
		refToUse = commit
	}
	return fmt.Sprintf("%s/blob/%s/%s",
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

		url = fmt.Sprintf("%s#L%d-L%d", url, startLine, endLine)
	}

	return url
}

func (e environment) IsCurrentEnvironment() bool {
	_, isExists := os.LookupEnv(githubWorkflowEnv)
	return isExists
}

func (e environment) Name() string {
	return "github"
}

func initPayload() (*GithubPayload, error) {
	var payload *GithubPayload

	payloadFile, err := os.ReadFile(os.Getenv(githubEventPath))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(payloadFile, &payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func GetCommits(payload *GithubPayload) []models.Commit {
	commits := make([]models.Commit, len(payload.Commits))
	for i, gc := range payload.Commits {
		commits[i] = models.Commit{
			Id:         gc.Id,
			Message:    gc.Message,
			CommitDate: gc.Timestamp,
			Url:        gc.Url,
			Author: models.Author{
				Email:    gc.Author.Email,
				Name:     gc.Author.Name,
				Username: gc.Author.Username,
			},
		}
	}

	return commits
}

func GetPipelinePaths(rootDir string) []string {
	paths := make([]string, 0)

	rootDirDepth := len(strings.Split(rootDir, "/"))

	filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == githubDir || strings.HasSuffix(path, workflowsDir) {
				return nil
			}
			if len(strings.Split(path, "/"))-rootDirDepth > 1 {
				return fs.SkipDir
			}
		}
		if strings.Contains(path, workflowsDir) {
			if filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml" {
				paths = append(paths, path)
			}
			return nil
		}
		return nil
	})

	return paths
}
