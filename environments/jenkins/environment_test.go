package jenkins

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
	jenkinsGithubMainFullEnvsFilePath    = "testdata/jenkins-github-main-full-env.json"
	jenkinsGithubMainNoGitEnvsFilePath   = "testdata/jenkins-github-main-no-.git-env.json"
	jenkinsGithubMainMinimalEnvsFilePath = "testdata/jenkins-github-main-minimal-env.json"
	testRepoPath                         = "/tmp/jenkins/repo"
	testRepoUrl                          = "https://github.com/test-organization/test-repo"
	testRepoCloneUrl                     = fmt.Sprintf("%s%s", testRepoUrl, ".git")
	testRepoCommit                       = "kcy8v2oazy4acuo1475rgtzg0nh403p23vw812lv"
	testdataPath                         = "../jenkins/testdata/repo"
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
			name:         "Jenkins GitHub main environment full env vars",
			envsFilePath: jenkinsGithubMainFullEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://test-jenkins.com:8080/",
				SCMApiUrl: githubApiUrl,
				Builder:   "Jenkins",
				LocalPath: testRepoPath,
				CommitSha: "kcy8v2oazy4acuo1475rgtzg0nh403p23vw812lv",
				Branch:    "main",
				Repository: models.Repository{
					Name:     "test-repo",
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Github,
				},
				Organization: models.Entity{
					Name: "test-organization",
				},
				Pipeline: models.Pipeline{
					Entity: models.Entity{
						Id:   "test-project",
						Name: "test-project",
					},
				},

				Job: models.Entity{
					Id:   "Run all",
					Name: "Run all",
				},
				Run: models.BuildRun{
					BuildId:     "4",
					BuildNumber: "4",
				},
				Runner: models.Runner{
					Id:           "master",
					Name:         "master",
					OS:           runtime.GOOS,
					Architecture: runtime.GOARCH,
				},
				PullRequest: models.PullRequest{
					SourceRef: models.Ref{
						Branch: "main",
					},
					TargetRef: models.Ref{
						Branch: "",
					},
				},
				PipelinePaths: []string{"/tmp/jenkins/repo/Jenkinsfile"},
				Environment:   enums.Jenkins,
				ScmId:         "8891c0db39f3064732cc1b4ac02c9b9f",
			},
			wantErr: false,
		},
		{
			name:         "Jenkins GitHub main environment minimal env vars",
			envsFilePath: jenkinsGithubMainMinimalEnvsFilePath,
			gitClient:    (&mocks.MockGitClient{}).SetRemoteUrl(testRepoCloneUrl).SetCommit(testRepoCommit).SetBranch("main"),
			want: &models.Configuration{
				Url:       "https://test-jenkins.com:8080/",
				SCMApiUrl: githubApiUrl,
				Builder:   "Jenkins",
				LocalPath: testRepoPath,
				CommitSha: "kcy8v2oazy4acuo1475rgtzg0nh403p23vw812lv",
				Branch:    "main",
				Repository: models.Repository{
					Name:     "test-repo",
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Github,
				},
				Organization: models.Entity{
					Name: "test-organization",
				},
				Pipeline: models.Pipeline{
					Entity: models.Entity{
						Id:   "test-project",
						Name: "test-project",
					},
				},

				Job: models.Entity{
					Id:   "Run all",
					Name: "Run all",
				},
				Run: models.BuildRun{
					BuildId:     "4",
					BuildNumber: "4",
				},
				Runner: models.Runner{
					Id:           "master",
					Name:         "master",
					OS:           runtime.GOOS,
					Architecture: runtime.GOARCH,
				},
				PullRequest: models.PullRequest{
					SourceRef: models.Ref{
						Branch: "main",
					},
					TargetRef: models.Ref{
						Branch: "",
					},
				},
				PipelinePaths: []string{"/tmp/jenkins/repo/Jenkinsfile"},
				Environment:   enums.Jenkins,
				ScmId:         "8891c0db39f3064732cc1b4ac02c9b9f",
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

func Test_environment_GetStepLink(t *testing.T) {
	tests := []struct {
		name         string
		envsFilePath string
		want         string
	}{
		{
			name:         "Jenkins environment",
			envsFilePath: jenkinsGithubMainFullEnvsFilePath,
			want:         "https://test-jenkins.com:8080/job/test-project/4/display/redirect",
		},
		{
			name:         "Not Jenkins environment",
			envsFilePath: "",
			want:         "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			if got := e.GetStepLink(); got != tt.want {
				t.Errorf("environment.GetStepLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_environment_GetBuildLink(t *testing.T) {
	tests := []struct {
		name         string
		envsFilePath string
		want         string
	}{
		{
			name:         "Jenkins environment",
			envsFilePath: jenkinsGithubMainFullEnvsFilePath,
			want:         "https://test-jenkins.com:8080/job/test-project/4/",
		},
		{
			name:         "Not Jenkins environment",
			envsFilePath: "",
			want:         "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			if got := e.GetBuildLink(); got != tt.want {
				t.Errorf("environment.GetBuildLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_environment_GetFileLineLink(t *testing.T) {
	type args struct {
		filename  string
		branch    string
		commit    string
		startLine int
		endLine   int
	}
	tests := []struct {
		name         string
		envsFilePath string
		args         args
		want         string
	}{
		{
			name:         "Jenkins environment",
			envsFilePath: jenkinsGithubMainFullEnvsFilePath,
			want:         "",
		},
		{
			name:         "Not Jenkins environment",
			envsFilePath: "",
			want:         "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			if got := e.GetFileLineLink(tt.args.filename, tt.args.branch, tt.args.commit, tt.args.startLine, tt.args.endLine); got != tt.want {
				t.Errorf("environment.GetFileLineLink() = %v, want %v", got, tt.want)
			}
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
			name:         "Jenkins environment",
			envsFilePath: jenkinsGithubMainFullEnvsFilePath,
			want:         true,
		},
		{
			name:         "Not Jenkins environment",
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

func Test_getRepositoryCloneURL(t *testing.T) {
	type args struct {
		repositoryPath string
	}
	tests := []struct {
		name         string
		envsFilePath string
		gitClient    *mocks.MockGitClient
		args         args
		want         string
		wantErr      bool
	}{
		{
			name:         "Clone url in env var with .git",
			envsFilePath: jenkinsGithubMainFullEnvsFilePath,
			args: args{
				repositoryPath: "",
			},
			want:    "https://github.com/test-organization/test-repo.git",
			wantErr: false,
		},
		{
			name:         "Clone url in env var without .git",
			envsFilePath: jenkinsGithubMainNoGitEnvsFilePath,
			args: args{
				repositoryPath: "",
			},
			want:    "https://github.com/test-organization/test-repo.git",
			wantErr: false,
		},
		{
			name:         "Clone url from repositoryPath with .git",
			envsFilePath: jenkinsGithubMainMinimalEnvsFilePath,
			gitClient:    (&mocks.MockGitClient{}).SetRemoteUrl("https://github.com/test-organization/test-repo.git"),
			args: args{
				repositoryPath: "",
			},
			want:    "https://github.com/test-organization/test-repo.git",
			wantErr: false,
		},
		{
			name:         "Clone url from repositoryPath without .git",
			envsFilePath: jenkinsGithubMainMinimalEnvsFilePath,
			gitClient:    (&mocks.MockGitClient{}).SetRemoteUrl("https://github.com/test-organization/test-repo"),
			args: args{
				repositoryPath: "",
			},
			want:    "https://github.com/test-organization/test-repo.git",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				prepareTest(t, tt.envsFilePath)
				setMockGitClient(t, tt.gitClient)
				got, err := getRepositoryCloneURL(tt.args.repositoryPath)
				if (err != nil) != tt.wantErr {
					t.Errorf("getRepositoryCloneURL() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("getRepositoryCloneURL() = %v, want %v", got, tt.want)
				}
			})
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
