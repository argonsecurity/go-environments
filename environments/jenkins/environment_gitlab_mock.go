package jenkins

import (
	"fmt"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/models"
)

var (
	MockGitlabOrgName   = "test-org"
	MockGitlabSubgroups = "sub1/sub2"
	MockGitlabRepoName  = "test-repo"
)

var mockGitlabConfiguration *models.Configuration

type EnvironmentGitlabMock struct{}

func (em *EnvironmentGitlabMock) GetConfiguration() (*models.Configuration, error) {
	if mockGitlabConfiguration == nil {
		if err := loadMockGitlabConfiguration(); err != nil {
			return nil, err
		}
	}
	return mockGitlabConfiguration, nil
}

func loadMockGitlabConfiguration() error {
	mockGitlabConfiguration = &models.Configuration{
		Url:             "https://server-gitlab.company.com",
		SCMApiUrl:       "https://server-gitlab.company.com",
		LocalPath:       "/path/to/repo",
		Branch:          "main",
		CommitSha:       "c6322vbd859aaew726d1e05ee1fc116365b90454",
		BeforeCommitSha: "0000000000000000000000000000000000000000",
		Repository: models.Repository{
			Name:     MockGitlabRepoName,
			CloneUrl: fmt.Sprintf("https://gitlab-server.company.com/test_user/%s", MockGitlabRepoName),
			Source:   enums.GitlabServer,
			Url:      fmt.Sprintf("https://gitlab-server.company.com/test_user/%s.git", MockGitlabRepoName),
		},
		Pipeline: models.Pipeline{
			Entity: models.Entity{
				Name: "test pipeline",
			},
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
				Branch: "main",
			},
			TargetRef: models.Ref{
				Branch: "",
			},
		},
		Builder: "Jenkins",
		Organization: models.Entity{
			Name: MockGitlabOrgName,
		},
		Pusher: models.Pusher{
			Entity: models.Entity{
				Name: "test user",
			},
			Username: "test_user",
		},
		PipelinePaths: []string{},
		Environment:   enums.Jenkins,
	}
	return nil
}

func (em *EnvironmentGitlabMock) GetBuildLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockGitlabOrgName, MockGitlabRepoName)
}

func (em *EnvironmentGitlabMock) GetStepLink() string {
	return fmt.Sprintf("http://localhost:8080/job/%s/job/%s/job/main/8/display/redirect", MockGitlabOrgName, MockGitlabRepoName)
}

func (em *EnvironmentGitlabMock) GetFileLineLink(filename string, ref string, line int) string {
	return ""
}

func (em *EnvironmentGitlabMock) IsCurrentEnvironment() bool {
	return true
}

func (em *EnvironmentGitlabMock) Name() string {
	return "jenkins"
}
