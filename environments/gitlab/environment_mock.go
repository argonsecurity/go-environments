package gitlab

import (
	"fmt"
	"net/url"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/models"
)

var (
	MockOrgName   = "mock-org"
	MockRepoName  = "mock-repo"
	MockSubGroups = "test-1/test-2"
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
		Url:             "https://gitlab.com",
		SCMApiUrl:       "https://gitlab.com",
		LocalPath:       "/path/to/gitlab/repo",
		CommitSha:       "3s32e4s818c6d1s5a0f585sf73112673a9bfcfc7",
		BeforeCommitSha: "0000000000000000000000000000000000000000",
		Branch:          "master",
		Run: models.BuildRun{
			BuildId:     "5131544634",
			BuildNumber: "mock run",
		},
		Pipeline: models.Pipeline{
			Entity: models.Entity{
				Id:   "373516881",
				Name: "mock pipeline",
			},
		},
		Runner: models.Runner{
			Id: "63824405",
			OS: "linux/arm64",
		},
		Repository: models.Repository{
			Id:       "12345678",
			Name:     MockRepoName,
			Url:      fmt.Sprintf("https://gitlab.com/%s/%s/%s", MockOrgName, MockSubGroups, MockRepoName),
			CloneUrl: fmt.Sprintf("https://gitlab-ci-token:[MASKED]@gitlab.com/%s/%s/%s.git", MockOrgName, MockSubGroups, MockRepoName),
			Source:   enums.Gitlab,
		},
		PullRequest: models.PullRequest{
			SourceRef: models.Ref{
				Branch: "",
			},
			TargetRef: models.Ref{
				Branch: "",
			},
		},
		Commits: []models.Commit{},
		Builder: "",
		Organization: models.Entity{
			Name: "",
		},
		Pusher: models.Pusher{
			Entity: models.Entity{
				Id:   "",
				Name: "",
			},
		},
		PipelinePaths: []string{},
		Environment:   enums.Gitlab,
	}

	return nil
}

func (em *EnvironmentMock) GetBuildLink() string {
	return fmt.Sprintf("https://gitlab.com/%s/%s/%s/-/pipelines/473114865", MockOrgName, MockSubGroups, MockRepoName)
}

func (em *EnvironmentMock) GetStepLink() string {
	return fmt.Sprintf("https://gitlab.com/%s/%s/%s/-/jobs/473114865", MockOrgName, MockSubGroups, MockRepoName)
}

func (em *EnvironmentMock) GetFileLineLink(filePath string, ref string, line int) string {
	return fmt.Sprintf("https://gitlab.com/%s/%s/%s/-/blob/%s/%s", MockOrgName, MockSubGroups, MockRepoName, ref, url.PathEscape(filePath))
}

func (em *EnvironmentMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentMock) Name() string {
	return "gitlab"
}
