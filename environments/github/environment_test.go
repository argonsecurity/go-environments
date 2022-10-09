package github

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/testutils"
	"github.com/argonsecurity/go-environments/models"
	"github.com/stretchr/testify/assert"
)

const (
	githubMainEnvsFilePath   = "testdata/github-workflows-main-env.json"
	githubPrEnvsFilePath     = "testdata/github-workflows-pr-env.json"
	githubServerEnvsFilePath = "testdata/github-server-workflows-main-env.json"
	testRepoPath             = "/tmp/github/repo"
	testRepoUrl              = "https://github.com/test-org/test-repo"
	testdataPath             = "../github/testdata/repo"

	testBranch   = "branch"
	testCommit   = "commit"
	testFilepath = "path/to/file"
)

var (
	testRepoCloneUrl = fmt.Sprintf("%s%s", testRepoUrl, ".git")
)

func Test_environment_GetConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		envsFilePath string
		want         *models.Configuration
		wantErr      bool
	}{
		{
			name:         "GitHub main configuration",
			envsFilePath: githubMainEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://github.com",
				SCMApiUrl: "https://api.github.com",
				LocalPath: testRepoPath,
				CommitSha: "2c6e3880fd94ddb7ef72d34e683cdc0c47bec6e6",
				Branch:    "refs/heads/main",
				Run: models.BuildRun{
					BuildId:     "3008488429",
					BuildNumber: "3",
				},
				Job: models.Entity{
					Id:   "test",
					Name: "test",
				},
				Pipeline: models.Entity{
					Id:   "test",
					Name: "test",
				},
				Runner: models.Runner{
					Id:           "3008488429",
					Name:         "Hosted Agent",
					OS:           "Linux",
					Architecture: runtime.GOARCH,
				},
				Repository: models.Repository{
					Id:       "507947722",
					Name:     "test-repo",
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
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
						Id:         "2c6e3880fd94ddb7ef72d34e683cdc0c47bec6e6",
						Message:    "Commit message",
						CommitDate: "2022-09-07T17:38:04+03:00",
						Url:        "https://github.com/test-org/test-repo/commit/2c6e3880fd94ddb7ef72d34e683cdc0c47bec6e6",
						Author: models.Author{
							Email:    "12345678+username123@users.noreply.github.com",
							Name:     "User Name",
							Username: "username123",
						},
					},
				},
				Builder: "Github Action",
				Organization: models.Entity{
					Name: "test-org",
				},
				Pusher: models.Pusher{
					Username: "username123",
					Entity: models.Entity{
						Id:   "19283746",
						Name: "username123",
					},
				},
				PipelinePaths: []string{
					filepath.Join(testRepoPath, ".github/workflows/first.yml"),
					filepath.Join(testRepoPath, ".github/workflows/second.yaml"),
				},
				Environment: enums.Github,
				ScmId:       "b30f418cdcc9970849d3d031de5df54f",
			},
			wantErr: false,
		},
		{
			name:         "GitHub PR configuration",
			envsFilePath: githubPrEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://github.com",
				SCMApiUrl: "https://api.github.com",
				LocalPath: testRepoPath,
				CommitSha: "mky2jknpc4fuz6qsn0vtouqwfjbno39itu0hifvs",
				Branch:    "test-branch",
				Run: models.BuildRun{
					BuildId:     "3014839969",
					BuildNumber: "6",
				},
				Job: models.Entity{
					Id:   "test",
					Name: "test",
				},
				Pipeline: models.Entity{
					Id:   "test",
					Name: "test",
				},
				Runner: models.Runner{
					Id:           "3014839969",
					Name:         "Hosted Agent",
					OS:           "Linux",
					Architecture: runtime.GOARCH,
				},
				Repository: models.Repository{
					Id:       "19283746",
					Name:     "test-repo",
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Github,
				},
				PullRequest: models.PullRequest{
					SourceRef: models.Ref{
						Branch: "test-branch",
					},
					TargetRef: models.Ref{
						Branch: "main",
					},
				},
				Commits: []models.Commit{},
				Builder: "Github Action",
				Organization: models.Entity{
					Name: "test-org",
				},
				Pusher: models.Pusher{
					Username: "username123",
					Entity: models.Entity{
						Id:   "99887766",
						Name: "username123",
					},
				},
				PipelinePaths: []string{
					filepath.Join(testRepoPath, ".github/workflows/first.yml"),
					filepath.Join(testRepoPath, ".github/workflows/second.yaml"),
				},
				Environment: enums.Github,
				ScmId:       "b30f418cdcc9970849d3d031de5df54f",
			},
			wantErr: false,
		},
		{
			name:         "GitHub Server main configuration",
			envsFilePath: githubServerEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://github.test.com",
				SCMApiUrl: "https://github.test.com/api/v3",
				LocalPath: testRepoPath,
				CommitSha: "2c6e3880fd94ddb7ef72d34e683cdc0c47bec6e6",
				Branch:    "refs/heads/main",
				Run: models.BuildRun{
					BuildId:     "3008488429",
					BuildNumber: "3",
				},
				Job: models.Entity{
					Id:   "test",
					Name: "test",
				},
				Pipeline: models.Entity{
					Id:   "test",
					Name: "test",
				},
				Runner: models.Runner{
					Id:           "3008488429",
					Name:         "Hosted Agent",
					OS:           "Linux",
					Architecture: runtime.GOARCH,
				},
				Repository: models.Repository{
					Id:       "507947722",
					Name:     "test-repo",
					Url:      "https://github.test.com/test-org/test-repo",
					CloneUrl: "https://github.com/test-org/test-repo.git",
					Source:   enums.GithubServer,
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
						Id:         "2c6e3880fd94ddb7ef72d34e683cdc0c47bec6e6",
						Message:    "Commit message",
						CommitDate: "2022-09-07T17:38:04+03:00",
						Url:        "https://github.com/test-org/test-repo/commit/2c6e3880fd94ddb7ef72d34e683cdc0c47bec6e6",
						Author: models.Author{
							Email:    "12345678+username123@users.noreply.github.com",
							Name:     "User Name",
							Username: "username123",
						},
					},
				},
				Builder: "Github Action",
				Organization: models.Entity{
					Name: "test-org",
				},
				Pusher: models.Pusher{
					Username: "username123",
					Entity: models.Entity{
						Id:   "19283746",
						Name: "username123",
					},
				},
				PipelinePaths: []string{
					filepath.Join(testRepoPath, ".github/workflows/first.yml"),
					filepath.Join(testRepoPath, ".github/workflows/second.yaml"),
				},
				Environment: enums.GithubServer,
				ScmId:       "b30f418cdcc9970849d3d031de5df54f",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
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
			name:         "GitHub environment",
			envsFilePath: githubMainEnvsFilePath,
			want:         "https://github.com/test-org/test-repo/actions/runs/3008488429",
		},
		{
			name:         "Not GitHub environment",
			envsFilePath: "",
			want:         "//actions/runs/",
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
			name:         "GitHub environment",
			envsFilePath: githubMainEnvsFilePath,
			want:         "https://github.com/test-org/test-repo/actions/runs/3008488429",
		},
		{
			name:         "Not GitHub environment",
			envsFilePath: "",
			want:         "//actions/runs/",
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
		filePath  string
		branch    string
		commit    string
		startLine int
		endLine   int
	}
	tests := []struct {
		name         string
		args         args
		envsFilePath string
		want         string
	}{
		{
			name: "No lines",
			args: args{
				filePath: testFilepath,
				branch:   "branchName",
			},
			envsFilePath: githubMainEnvsFilePath,
			want:         "https://github.com/test-org/test-repo/blob/branchName/path/to/file",
		},
		{
			name: "Same line",
			args: args{
				filePath:  testFilepath,
				commit:    "1a70bx6328bad78d919dca422d1as1g1ec97c5f6",
				startLine: 1,
				endLine:   1,
			},
			envsFilePath: githubMainEnvsFilePath,
			want:         "https://github.com/test-org/test-repo/blob/1a70bx6328bad78d919dca422d1as1g1ec97c5f6/path/to/file#L1-L1",
		},
		{
			name: "Different lines",
			args: args{
				filePath:  "path/to/file",
				commit:    "1a70bx6328bad78d919dca422d1as1g1ec97c5f6",
				startLine: 1,
				endLine:   2,
			},
			envsFilePath: githubMainEnvsFilePath,
			want:         "https://github.com/test-org/test-repo/blob/1a70bx6328bad78d919dca422d1as1g1ec97c5f6/path/to/file#L1-L2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			if got := e.GetFileLineLink(tt.args.filePath, tt.args.branch, tt.args.commit, tt.args.startLine, tt.args.endLine); got != tt.want {
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
			name:         "GitHub main environment",
			envsFilePath: githubMainEnvsFilePath,
			want:         true,
		},
		{
			name:         "GitHub pr environment",
			envsFilePath: githubPrEnvsFilePath,
			want:         true,
		},
		{
			name:         "Not GitHub environment",
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

func TestGetFileLink(t *testing.T) {
	type args struct {
		repositoryURL string
		filename      string
		branch        string
		commit        string
		startLine     int
		endLine       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No line numbers",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testFilepath,
				branch:        testBranch,
			},
			want: "https://github.com/test-org/test-repo/blob/branch/path/to/file",
		},
		{
			name: "Same line",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testFilepath,
				branch:        testBranch,
				startLine:     1,
				endLine:       1,
			},
			want: "https://github.com/test-org/test-repo/blob/branch/path/to/file#L1-L1",
		},
		{
			name: "Different lines",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testFilepath,
				branch:        testBranch,
				startLine:     1,
				endLine:       2,
			},
			want: "https://github.com/test-org/test-repo/blob/branch/path/to/file#L1-L2",
		},
		{
			name: "With commit",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testFilepath,
				commit:        testCommit,
				startLine:     1,
				endLine:       2,
			},
			want: "https://github.com/test-org/test-repo/blob/commit/path/to/file#L1-L2",
		},
		{
			name: "With commit and branch",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testFilepath,
				commit:        testCommit,
				branch:        testBranch,
				startLine:     1,
				endLine:       2,
			},
			want: "https://github.com/test-org/test-repo/blob/commit/path/to/file#L1-L2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileLink(tt.args.repositoryURL, tt.args.filename, tt.args.branch, tt.args.commit, tt.args.startLine, tt.args.endLine); got != tt.want {
				t.Errorf("GetFileLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
