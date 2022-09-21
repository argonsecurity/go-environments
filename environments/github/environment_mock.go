package github

import (
	"fmt"
	"net/url"

	"github.com/argonsecurity/go-utils/environments/enums"
	"github.com/argonsecurity/go-utils/environments/models"
)

var (
	MockOrgName  = "test-org"
	MockRepoName = "test-repo"
)

var mockConfiguration *models.Configuration

type EnvironmentMock struct{}

func (em *EnvironmentMock) GetConfiguration() (*models.Configuration, error) {
	if mockConfiguration == nil {
		if err := loadMockConfiguration(); err != nil {
			return nil, err
		}
	}
	return mockConfiguration, nil
}

func loadMockConfiguration() error {
	mockConfiguration = &models.Configuration{
		Url:       "https://github.com",
		SCMApiUrl: "https://api.github.com",
		LocalPath: "/path/to/repo",
		CommitSha: "1bhf303744315bc0a5e9945f9f4df4b3f3dab272",
		Branch:    "refs/heads/master",
		Run: models.Entity{
			Id: "2837115383",
		},
		Pipeline: models.Entity{
			Name: "test-pipeline",
		},
		Runner: models.Runner{
			Id:   "2837115383",
			Name: "Hosted Agent",
			OS:   "Linux",
		},
		Repository: models.Repository{
			Id:       "410143034",
			Name:     MockRepoName,
			Url:      fmt.Sprintf("https://github.com/%s/%s", MockOrgName, MockRepoName),
			CloneUrl: "",
			Source:   enums.Github,
		},
		PullRequest: models.PullRequest{
			SourceRef: models.Ref{
				Branch: "",
			},
			TargetRef: models.Ref{
				Branch: "",
			},
		},
		Commits: []models.Commit{
			{
				Id:         "4499b33d2sa5755a2e85ffa617bsf814e556156c",
				Message:    "Some mock commit message",
				CommitDate: "2022-02-13T16:00:11+02:00",
				Url:        fmt.Sprintf("https://github.com/%s/%s/commit/4499b33d2sa5755a2e85ffa617bsf814e556156c", MockOrgName, MockRepoName),
				Author: models.Author{
					Email:    "67327471+mock.mail@users.noreply.github.com",
					Name:     "Mock User",
					Username: "MockUser123",
				},
			},
		},
		Builder: "Github Action",
		Organization: models.Entity{
			Name: MockOrgName,
		},
		Pusher: models.Pusher{
			Entity: models.Entity{
				Id:   "67327471",
				Name: "MockUser123",
			},
		},
		PipelinePaths: []string{},
		Environment:   enums.Github,
	}
	return nil
}

func (em *EnvironmentMock) GetBuildLink() string {
	return fmt.Sprintf("https://github.com/%s/%s/actions/runs/1837111563", MockOrgName, MockRepoName)
}

func (em *EnvironmentMock) GetStepLink() string {
	return fmt.Sprintf("https://github.com/%s/%s/actions/runs/1837111563", MockOrgName, MockRepoName)
}

func (em *EnvironmentMock) GetFileLineLink(filePath string, ref string, line int) string {
	return fmt.Sprintf("https://github.com/%s/%s/blob/%s/%s", MockOrgName, MockRepoName, ref, url.PathEscape(filePath))
}

func (em *EnvironmentMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentMock) Name() string {
	return "github"
}
