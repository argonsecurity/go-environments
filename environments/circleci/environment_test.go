package circleci

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/testutils"
	"github.com/argonsecurity/go-environments/environments/testutils/mocks"
	"github.com/argonsecurity/go-environments/environments/utils/git"
	"github.com/argonsecurity/go-environments/models"
	"github.com/stretchr/testify/assert"
)

var (
	circleciGithubMainFullEnvsFilePath = "testdata/circleci-github-main-env.json"
	testRepoPath                       = "/tmp/circleci/repo"
	testRepoUrl                        = "https://github.com/test-organization/test-repo"
	testRepoCloneUrl                   = fmt.Sprintf("%s%s", testRepoUrl, ".git")
	testdataPath                       = "../circleci/testdata/repo"
)

func Test_environment_GetConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		envsFilePath string
		gitClient    *mocks.MockGitClient
		want         *models.Configuration
		wantErr      bool
	}{
		{
			name:         "CircleCi GitHub main environment env vars",
			envsFilePath: circleciGithubMainFullEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://app.circleci.com",
				SCMApiUrl: githubApiUrl,
				Builder:   "CircleCi",
				LocalPath: "https://github.com/test-organization/test-repo.git",
				CommitSha: "kcy8v2oazy4acuo1475rgtzg0nh403p23vw812lv",
				Branch:    "main",
				Repository: models.Repository{
					Name:     "test-repo",
					FullName: "test-organization/test-repo",
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Github,
				},
				Organization: models.Entity{
					Name: "test-organization",
				},
				Run: models.BuildRun{
					BuildId:     "4",
					BuildNumber: "4",
				},
				Runner: models.Runner{

					OS:           runtime.GOOS,
					Architecture: runtime.GOARCH,
				},
				PullRequest: models.PullRequest{
					Id: "49",
				},
				Environment: enums.CircleCi,
				ScmId:       "8891c0db39f3064732cc1b4ac02c9b9f",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			setMockGitClient(t, tt.gitClient)
			got, err := e.GetConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("environment.GetConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_environment_IsCurrentEnvironment(t *testing.T) {
	tests := []struct {
		name         string
		envsFilePath string
		want         bool
	}{
		{
			name:         "CircelCi environment",
			envsFilePath: circleciGithubMainFullEnvsFilePath,
			want:         true,
		},
		{
			name:         "Not CircelCi environment",
			envsFilePath: "",
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			if got := e.IsCurrentEnvironment(); got != tt.want {
				t.Errorf("environment.IsCurrentEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func prepareTest(t *testing.T, envsFilePath string) environment {
	e := environment{}
	configuration = nil
	testRepoCleanup := testutils.PrepareTestGitRepository(testRepoPath, testRepoCloneUrl, testdataPath)
	t.Cleanup(testRepoCleanup)
	envCleanup := testutils.SetEnvsFromFile(envsFilePath)
	t.Cleanup(envCleanup)
	return e
}

func setMockGitClient(t *testing.T, gitClient *mocks.MockGitClient) {
	originalClient := git.GlobalGitClient
	mockGitClient := gitClient
	if mockGitClient == nil {
		mockGitClient = &mocks.MockGitClient{}
	}
	mocks.SetGitMock(mockGitClient)
	t.Cleanup(func() {
		git.GlobalGitClient = originalClient
	})
}
