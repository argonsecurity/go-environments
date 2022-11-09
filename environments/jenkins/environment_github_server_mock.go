package jenkins

import (
	"fmt"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/models"
)

var (
	MockGithubServerOrgName  = "mock-org"
	MockGithubServerRepoName = "mock-repo"
)

var mockGithubServerConfiguration *models.Configuration

type EnvironmentGithubServerMock struct{}

func (em *EnvironmentGithubServerMock) GetConfiguration() (*models.Configuration, error) {
	if mockGithubServerConfiguration == nil {
		if err := loadMockGithubServerConfiguration(); err != nil {
			return nil, err
		}
	}
	return mockGithubServerConfiguration, nil
}

func loadMockGithubServerConfiguration() error {
	mockGithubServerConfiguration = &models.Configuration{
		Url:       "https://github.server.com",
		SCMApiUrl: githubApiUrl,
		LocalPath: "/path/to/repo",
		Branch:    "main",
		CommitSha: "ab1272140f7c845cb8ea3d18r08174s546fb2c75",
		Repository: models.Repository{
			Name:     MockGithubServerRepoName,
			CloneUrl: fmt.Sprintf("git@github.server.com:%s/%s.git", MockGithubServerOrgName, MockGithubServerRepoName),
			Source:   enums.GithubServer,
			Url:      fmt.Sprintf("https://github.server.com/api/v3/%s/%s", MockGithubServerOrgName, MockGithubServerRepoName),
		},
		Pipeline: models.Pipeline{
			Entity: models.Entity{
				Id:   "1",
				Name: "mock-pipeline",
			},
		},
		Run: models.BuildRun{
			BuildId:     "jenkins-mock-pipeline-1",
			BuildNumber: "1",
		},
		Runner: models.Runner{
			Id:   "built_in",
			Name: "built_in",
			OS:   "darwin",
		},
		PullRequest: models.PullRequest{
			SourceRef: models.Ref{
				Branch: "main",
			},
			TargetRef: models.Ref{
				Branch: "",
			},
		},
		Builder: "Jenkins",
		Organization: models.Entity{
			Name: MockGithubServerOrgName,
		},
		PipelinePaths: []string{},
		Environment:   enums.Jenkins,
	}
	return nil
}

func (em *EnvironmentGithubServerMock) GetBuildLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockGithubServerOrgName, MockGithubServerRepoName)
}

func (em *EnvironmentGithubServerMock) GetStepLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockGithubServerOrgName, MockGithubServerRepoName)
}

func (e EnvironmentGithubServerMock) GetFileLink(filename string, branch string, commit string) string {
	return ""
}

func (e EnvironmentGithubServerMock) GetFileLineLink(filename string, branch string, commit string, startLine int, endLine int) string {
	return ""
}

func (em *EnvironmentGithubServerMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentGithubServerMock) Name() string {
	return "jenkins"
}
