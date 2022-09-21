package jenkins

import (
	"fmt"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/models"
)

var (
	MockBitbucketWorkspace = "test-org"
	MockBitbucketRepo      = "test-repo"
)

var mockBitbucketConfiguration *models.Configuration

type EnvironmentBitbucketMock struct{}

func (em *EnvironmentBitbucketMock) GetConfiguration() (*models.Configuration, error) {
	if mockBitbucketConfiguration == nil {
		if err := loadMockBitbucketConfiguration(); err != nil {
			return nil, err
		}
	}
	return mockBitbucketConfiguration, nil
}

func loadMockBitbucketConfiguration() error {
	mockBitbucketConfiguration = &models.Configuration{
		Url:             "https://bitbucket.org",
		SCMApiUrl:       bitbucketApiUrl,
		LocalPath:       "/path/to/repo",
		Branch:          "branch",
		CommitSha:       "c6322vbd859aaew726d1e05ee1fc116c65b9e454",
		BeforeCommitSha: "0000000000000000000000000000000000000000",
		Repository: models.Repository{
			Id:       "{fw7cs8e3-11a6-483c-bb50-765555349f3d}",
			Name:     MockBitbucketRepo,
			CloneUrl: fmt.Sprintf("https://bitbucket.org/%s/%s", MockBitbucketWorkspace, MockBitbucketRepo),
			Source:   enums.Bitbucket,
			Url:      fmt.Sprintf("https://bitbucket.org/%s/%s.git", MockBitbucketWorkspace, MockBitbucketRepo),
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
			Name: MockBitbucketWorkspace,
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

func (em *EnvironmentBitbucketMock) GetBuildLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockBitbucketWorkspace, MockBitbucketRepo)
}

func (em *EnvironmentBitbucketMock) GetStepLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockBitbucketWorkspace, MockBitbucketRepo)
}

func (em *EnvironmentBitbucketMock) GetFileLineLink(filename string, ref string, line int) string {
	return ""
}

func (em *EnvironmentBitbucketMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentBitbucketMock) Name() string {
	return "jenkins"
}
