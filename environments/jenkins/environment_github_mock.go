package jenkins

import (
	"fmt"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/models"
)

var (
	MockGithubOrgName  = "mock-org"
	MockGithubRepoName = "mock-repo"
)

var mockGithubConfiguration *models.Configuration

type EnvironmentGithubMock struct{}

func (em *EnvironmentGithubMock) GetConfiguration() (*models.Configuration, error) {
	if mockGithubConfiguration == nil {
		if err := loadMockGithubConfiguration(); err != nil {
			return nil, err
		}
	}
	return mockGithubConfiguration, nil
}

func loadMockGithubConfiguration() error {
	mockGithubConfiguration = &models.Configuration{
		Url:       "https://github.com",
		SCMApiUrl: githubApiUrl,
		LocalPath: "/path/to/repo",
		Branch:    "main",
		CommitSha: "ab1272140f7c845cb8ea3d18r08174s546fb2c75",
		Repository: models.Repository{
			Name:     MockGithubRepoName,
			CloneUrl: fmt.Sprintf("git@github.com:%s/%s.git", MockGithubOrgName, MockGithubRepoName),
			Source:   enums.Github,
			Url:      fmt.Sprintf("https://api.github.com/%s/%s", MockGithubOrgName, MockGithubRepoName),
		},
		Pipeline: models.Entity{
			Name: "mock pipeline",
		},
		Run: models.BuildRun{
			BuildId:     "1",
			BuildNumber: "jenkins-mock-pipeline-1",
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
			Name: MockGithubOrgName,
		},
		PipelinePaths: []string{},
		Environment:   enums.Jenkins,
	}
	return nil
}

func (em *EnvironmentGithubMock) GetBuildLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockGithubOrgName, MockGithubRepoName)
}

func (em *EnvironmentGithubMock) GetStepLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockGithubOrgName, MockGithubRepoName)
}

func (em *EnvironmentGithubMock) GetFileLineLink(filename string, ref string, line int) string {
	return ""
}

func (em *EnvironmentGithubMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentGithubMock) Name() string {
	return "jenkins"
}
