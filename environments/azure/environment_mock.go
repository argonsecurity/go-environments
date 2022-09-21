package azure

import (
	"fmt"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/models"
)

var (
	MockCollectionName = "test-collection"
	MockProjectName    = "test-project"
	MockRepoName       = "test-repo"
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
		Url:       fmt.Sprintf("https://dev.azure.com/%s/", MockCollectionName),
		SCMApiUrl: "https://dev.azure.com/",
		Branch:    "refs/heads/main",
		ProjectId: "75be3c85-72d8-45cb-92fb-633bc4bs6e23",
		CommitSha: "9c2dd65bxc45889d0y0c27fbdaa8cbcb5gx9h1a4",
		Repository: models.Repository{
			Id:       "519s4a2d-71c2-4966-bfc6-05a03601e443",
			Name:     MockRepoName,
			Url:      fmt.Sprintf("https://%s@dev.azure.com/%s/%s/_git/%s", MockCollectionName, MockCollectionName, MockProjectName, MockRepoName),
			CloneUrl: fmt.Sprintf("https://%s@dev.azure.com/%s/%s/_git/%s", MockCollectionName, MockCollectionName, MockProjectName, MockRepoName),
			Source:   enums.Azure,
		},
		Pusher: models.Pusher{
			Username: "user",
			Email:    "user@mock.com",
		},
		Pipeline: models.Entity{
			Id:   "752e0c81-72s8-44eb-93cb-1d3b5gbd1acc-63",
			Name: MockRepoName,
		},
		Run: models.BuildRun{
			BuildId: "3557",
		},
		Runner: models.Runner{
			Id:           "1",
			Name:         "",
			OS:           "Linux",
			Distribution: "ubuntu20",
			Architecture: "X64",
		},
		PullRequest: models.PullRequest{
			Id: "",
			SourceRef: models.Ref{
				Branch: "",
			},
			TargetRef: models.Ref{
				Branch: "",
			},
		},
		PipelinePaths: []string{"/path/to/pipeline"},
		Environment:   enums.Azure,
	}
	return nil
}

// GetBuildLink get a link to the current build
func (em *EnvironmentMock) GetBuildLink() string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/_build?definitionId=69&_a=summary", MockCollectionName, MockProjectName)
}

// GetStepLink get a link to the current step
func (em *EnvironmentMock) GetStepLink() string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/_build/results?buildId=1&view=logs&j=1&t=1", MockCollectionName, MockProjectName)
}

// GetFileLineLink get a link to a file line
func (em *EnvironmentMock) GetFileLineLink(filename string, ref string, line int) string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/", MockCollectionName, MockProjectName)
}

// IsCurrentEnvironment detects if the runtime environment matches the object
func (em *EnvironmentMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentMock) Name() string {
	return "azure"
}
