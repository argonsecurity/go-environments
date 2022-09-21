package bitbucket

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/models"
)

var (
	MockRepositoryName = "owner/repo"
	MockRunnerId       = "63824405"
	MockStepId         = "63824405"
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
		Url:       "https://bitbucket.org",
		SCMApiUrl: "https://api.bitbucket.org/2.0",
		LocalPath: "/path/to/bitbucket/repo",
		CommitSha: "3s32e4s818c6d1s5a0f585sf73112673a9bfcfc7",
		Branch:    "master",
		Run: models.BuildRun{
			BuildId: "5131544634",
		},
		Pipeline: models.Entity{
			Id:   "373516881",
			Name: "repo",
		},
		Runner: models.Runner{
			OS:           "linux",
			Architecture: "amd64",
		},
		Repository: models.Repository{
			Id:       "12345678",
			Name:     MockRepositoryName,
			Url:      fmt.Sprintf("https://bitbucket.org/%s", MockRepositoryName),
			CloneUrl: fmt.Sprintf("https://bitbucket.org/%s.git", MockRepositoryName),
			Source:   enums.Bitbucket,
		},
		Commits: []models.Commit{},
		Builder: "",
		Organization: models.Entity{
			Name: strings.Split(MockRepositoryName, "/")[0],
		},
		Pusher: models.Pusher{
			Entity: models.Entity{
				Id:   "",
				Name: "",
			},
		},
		PipelinePaths: []string{},
		Environment:   enums.Bitbucket,
	}

	return nil
}

func (em *EnvironmentMock) GetBuildLink() string {
	return fmt.Sprintf("https://bitbucket.org/%s/pipelines/%s", MockRepositoryName, MockRunnerId)
}

func (em *EnvironmentMock) GetStepLink() string {
	return fmt.Sprintf("https://bitbucket.org/%s/results/%s/steps/%s", MockRepositoryName, MockRunnerId, url.PathEscape(MockStepId))
}

func (em *EnvironmentMock) GetFileLineLink(filePath string, ref string, line int) string {
	return fmt.Sprintf("https:///bitbucket.org/%s/src/%s/%s#lines%d", MockRepositoryName, ref, filePath, line)
}

func (em *EnvironmentMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentMock) Name() string {
	return "bitbucket"
}
