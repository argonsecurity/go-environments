package jenkins

import (
	"fmt"

	"github.com/argonsecurity/go-utils/environments/enums"
	"github.com/argonsecurity/go-utils/environments/models"
)

var (
	MockBitbucketServerWorkspace = "test-org"
	MockBitbucketServerRepo      = "test-repo"
)

var mockBitbucketServerConfiguration *models.Configuration

type EnvironmentBitbucketServerMock struct{}

func (em *EnvironmentBitbucketServerMock) GetConfiguration() (*models.Configuration, error) {
	if mockBitbucketServerConfiguration == nil {
		if err := loadMockBitbucketServerConfiguration(); err != nil {
			return nil, err
		}
	}
	return mockBitbucketServerConfiguration, nil
}

func loadMockBitbucketServerConfiguration() error {
	mockBitbucketServerConfiguration = &models.Configuration{
		Url:             "https://staging-bitbucket.org",
		SCMApiUrl:       bitbucketApiUrl,
		LocalPath:       "/path/to/repo",
		Branch:          "branch",
		CommitSha:       "c6322vbd859aaew726d1e05ee1fc116c65b9e454",
		BeforeCommitSha: "0000000000000000000000000000000000000000",
		Repository: models.Repository{
			Id:       "{fw7cs8e3-11a6-483c-bb50-765555349f3d}",
			Name:     MockBitbucketServerRepo,
			CloneUrl: fmt.Sprintf("https://staging-bitbucket.org/%s/%s", MockBitbucketServerWorkspace, MockBitbucketServerRepo),
			Source:   enums.BitbucketServer,
			Url:      fmt.Sprintf("https://staging-bitbucket.org/%s/%s.git", MockBitbucketServerWorkspace, MockBitbucketServerRepo),
		},
		Run: models.BuildRun{
			BuildId:     "4",
			BuildNumber: "jenkins-test-pipeline-main-4",
		},
		Runner: models.Runner{
			Id:   "built_in",
			Name: "built_in",
			OS:   "darwin",
		},
		PullRequest: models.PullRequest{
			SourceRef: models.Ref{
				Branch: "branch-test",
				Sha:    "14a3f0b44b8s",
			},
			TargetRef: models.Ref{
				Branch: "main",
				Sha:    "74c3f0a4dv88",
			},
		},
		Builder: "Jenkins",
		Organization: models.Entity{
			Name: MockBitbucketServerWorkspace,
		},
		Pusher: models.Pusher{
			Entity: models.Entity{
				Id: "{f9638c63-sf38-fftb-bedf-4bw0ta1y585g}",
			},
			Username: "test_user",
		},
		PipelinePaths: []string{
			"/path/to/repo/Jenkinsfile",
		},
		Environment: enums.Jenkins,
	}
	return nil
}

func (em *EnvironmentBitbucketServerMock) GetBuildLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockBitbucketServerWorkspace, MockBitbucketServerRepo)
}

func (em *EnvironmentBitbucketServerMock) GetStepLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockBitbucketServerWorkspace, MockBitbucketServerRepo)
}

func (em *EnvironmentBitbucketServerMock) GetFileLineLink(filename string, ref string, line int) string {
	return ""
}

func (em *EnvironmentBitbucketServerMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentBitbucketServerMock) Name() string {
	return "jenkins"
}
