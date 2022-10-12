package gitlab

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/testutils"
	"github.com/argonsecurity/go-environments/models"
	"github.com/stretchr/testify/assert"
)

const (
	gitlabMainEnvsFilePath       = "testdata/gitlab-ci-main-env.json"
	gitlabServerMainEnvsFilePath = "testdata/gitlab-server-ci-main-env.json"
	gitlabPrEnvsFilePath         = "testdata/gitlab-ci-pr-env.json"
	testRepoPath                 = "/tmp/gitlab/repo"
	testRepoUrl                  = "https://gitlab.com/test-group/test-sub-group/test-project"
	testdataPath                 = "../gitlab/testdata/repo"

	testCommit = "commit"
	testBranch = "branch"
	testPath   = "path/to/file"
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
			name:         "GitLab main configuration",
			envsFilePath: gitlabMainEnvsFilePath,
			want: &models.Configuration{
				Url:             "https://gitlab.com",
				SCMApiUrl:       "https://gitlab.com",
				LocalPath:       testRepoPath,
				Branch:          "main",
				CommitSha:       "3ufl0xuicz460no9xck5j3xyyvk9w8m4j7bwr3ta",
				BeforeCommitSha: "0000000000000000000000000000000000000000",
				Organization: models.Entity{
					Name: "test-group",
				},
				Repository: models.Repository{
					Id:       "88227733",
					Name:     "test-project",
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Gitlab,
				},
				Pipeline: models.Entity{
					Id:   "109238745",
					Name: "test-project",
				},
				Job: models.Entity{
					Id:   "deploy-main",
					Name: "deploy-main",
				},
				Runner: models.Runner{
					Id:           "13243546",
					Name:         "2-green.shared.runners-manager.gitlab.com/default",
					OS:           "linux/amd64",
					Architecture: runtime.GOARCH,
				},
				Run: models.BuildRun{
					BuildId: "3210743970",
				},
				Pusher: models.Pusher{
					Username: "User Name",
				},
				PullRequest: models.PullRequest{
					Id: "",
					SourceRef: models.Ref{
						Sha:    "",
						Branch: "",
					},
					TargetRef: models.Ref{
						Sha:    "",
						Branch: "",
					},
				},
				PipelinePaths: []string{"/tmp/gitlab/repo/.gitlab-ci.yml", "/tmp/gitlab/repo/.gitlab-ci.yaml"},
				Environment:   enums.Gitlab,
				ScmId:         "fb240c83d76e50991d7470048e98058a",
			},
			wantErr: false,
		},
		{
			name:         "GitLab pr configuration",
			envsFilePath: gitlabPrEnvsFilePath,
			want: &models.Configuration{
				Url:             "https://gitlab.com",
				SCMApiUrl:       "https://gitlab.com",
				LocalPath:       testRepoPath,
				Branch:          "test-branch",
				CommitSha:       "l1lv78gwsiwq1j1pgyx27ky7eiqjs84r2oa294j4",
				BeforeCommitSha: "0000000000000000000000000000000000000000",
				Organization: models.Entity{
					Name: "test-group",
				},
				Repository: models.Repository{
					Id:       "12345678",
					Name:     "test-project",
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Gitlab,
				},
				Pipeline: models.Entity{
					Id:   "840375028",
					Name: "test-project",
				},
				Job: models.Entity{
					Id:   "deploy-branch",
					Name: "deploy-branch",
				},
				Runner: models.Runner{
					Id:           "98472957",
					Name:         "2-green.shared.runners-manager.gitlab.com/default",
					OS:           "linux/amd64",
					Architecture: runtime.GOARCH,
				},
				Run: models.BuildRun{
					BuildId: "5510622136",
				},
				Pusher: models.Pusher{
					Username: "User Name",
				},
				PullRequest: models.PullRequest{
					Id: "473847937",
					SourceRef: models.Ref{
						Sha:    "",
						Branch: "test-branch",
					},
					TargetRef: models.Ref{
						Sha:    "",
						Branch: "main",
					},
				},
				PipelinePaths: []string{"/tmp/gitlab/repo/.gitlab-ci.yml", "/tmp/gitlab/repo/.gitlab-ci.yaml"},
				Environment:   enums.Gitlab,
				ScmId:         "fb240c83d76e50991d7470048e98058a",
			},
			wantErr: false,
		},
		{
			name:         "GitLab Server main configuration",
			envsFilePath: gitlabServerMainEnvsFilePath,
			want: &models.Configuration{
				Url:             "https://gitlab.test.com",
				SCMApiUrl:       "https://gitlab.test.com",
				LocalPath:       testRepoPath,
				Branch:          "main",
				CommitSha:       "3ufl0xuicz460no9xck5j3xyyvk9w8m4j7bwr3ta",
				BeforeCommitSha: "0000000000000000000000000000000000000000",
				Organization: models.Entity{
					Name: "test-group",
				},
				Repository: models.Repository{
					Id:       "88227733",
					Name:     "test-project",
					Url:      "https://gitlab.test.com/test-group/test-sub-group/test-project",
					CloneUrl: "https://gitlab.com/test-group/test-sub-group/test-project.git",
					Source:   enums.GitlabServer,
				},
				Pipeline: models.Entity{
					Id:   "109238745",
					Name: "test-project",
				},
				Job: models.Entity{
					Id:   "deploy-main",
					Name: "deploy-main",
				},
				Runner: models.Runner{
					Id:           "13243546",
					Name:         "2-green.shared.runners-manager.gitlab.com/default",
					OS:           "linux/amd64",
					Architecture: runtime.GOARCH,
				},
				Run: models.BuildRun{
					BuildId: "3210743970",
				},
				PullRequest: models.PullRequest{
					Id: "",
					SourceRef: models.Ref{
						Sha:    "",
						Branch: "",
					},
					TargetRef: models.Ref{
						Sha:    "",
						Branch: "",
					},
				},
				Pusher: models.Pusher{
					Username: "User Name",
				},
				PipelinePaths: []string{"/tmp/gitlab/repo/.gitlab-ci.yml", "/tmp/gitlab/repo/.gitlab-ci.yaml"},
				Environment:   enums.GitlabServer,
				ScmId:         "fb240c83d76e50991d7470048e98058a",
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
			name:         "GitLab environment",
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/jobs/3210743970",
		},
		{
			name:         "Not GitLab environment",
			envsFilePath: "",
			want:         "///-/jobs/",
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
			name:         "GitLab environment",
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/pipelines/109238745",
		},
		{
			name:         "Not GitLab environment",
			envsFilePath: "",
			want:         "///-/pipelines/",
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

func Test_environment_GetFileLink(t *testing.T) {
	type args struct {
		filename string
		branch   string
		commit   string
	}
	tests := []struct {
		name         string
		envsFilePath string
		args         args
		want         string
	}{
		{
			name: "File from branch",
			args: args{
				filename: testPath,
				branch:   testBranch,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/branch/path/to/file",
		},
		{
			name: "File from commit",
			args: args{
				filename: testPath,
				commit:   testCommit,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/commit/path/to/file",
		},
		{
			name: "Empty file path",
			args: args{
				filename: "",
				commit:   testCommit,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/commit/",
		},
		{
			name: "Empty ref",
			args: args{
				filename: testPath,
				branch:   "",
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob//path/to/file",
		},
		{
			name: "Not GitLab environment",
			args: args{
				filename: testPath,
				branch:   testBranch,
			},
			envsFilePath: "",
			want:         "/-/blob/branch/path/to/file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			if got := e.GetFileLink(tt.args.filename, tt.args.branch, tt.args.commit); got != tt.want {
				t.Errorf("environment.GetFileLink() = %v, want %v", got, tt.want)
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
			name: "File from branch",
			args: args{
				filename:  testPath,
				branch:    testBranch,
				startLine: 1,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/branch/path/to/file#L1-1",
		},
		{
			name: "File from branch with line number 0",
			args: args{
				filename:  testPath,
				branch:    testBranch,
				startLine: 0,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/branch/path/to/file",
		},
		{
			name: "File from commit",
			args: args{
				filename:  testPath,
				commit:    testCommit,
				startLine: 1,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/commit/path/to/file#L1-1",
		},
		{
			name: "Empty file path",
			args: args{
				filename:  "",
				commit:    testCommit,
				startLine: 1,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/commit/#L1-1",
		},
		{
			name: "Empty ref",
			args: args{
				filename:  testPath,
				branch:    "",
				startLine: 1,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob//path/to/file#L1-1",
		},
		{
			name: "Not GitLab environment",
			args: args{
				filename:  testPath,
				branch:    testBranch,
				startLine: 1,
			},
			envsFilePath: "",
			want:         "/-/blob/branch/path/to/file#L1-1",
		},
		{
			name: "Different line numbers",
			args: args{
				filename:  testPath,
				branch:    testBranch,
				startLine: 1,
				endLine:   2,
			},
			envsFilePath: gitlabMainEnvsFilePath,
			want:         "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/branch/path/to/file#L1-2",
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

func TestGetFileLink(t *testing.T) {
	type args struct {
		repositoryURL string
		filename      string
		branch        string
		commit        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "With branch",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				branch:        testBranch,
			},
			want: "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/branch/path/to/file",
		},
		{
			name: "With commit",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				commit:        testCommit,
			},
			want: "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/commit/path/to/file",
		},
		{
			name: "With commit and branch",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				branch:        testBranch,
				commit:        testCommit,
			},
			want: "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/commit/path/to/file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileLink(tt.args.repositoryURL, tt.args.filename, tt.args.branch, tt.args.commit); got != tt.want {
				t.Errorf("GetFileLineLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileLineLink(t *testing.T) {
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
			name: "No lines",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				branch:        testBranch,
			},
			want: "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/branch/path/to/file",
		},
		{
			name: "One line",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				branch:        testBranch,
				startLine:     1,
			},
			want: "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/branch/path/to/file#L1-1",
		},
		{
			name: "Different lines",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				branch:        testBranch,
				startLine:     1,
				endLine:       2,
			},
			want: "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/branch/path/to/file#L1-2",
		},
		{
			name: "With commit",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				commit:        testCommit,
				startLine:     1,
				endLine:       2,
			},
			want: "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/commit/path/to/file#L1-2",
		},
		{
			name: "With commit and branch",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				branch:        testBranch,
				commit:        testCommit,
				startLine:     1,
				endLine:       2,
			},
			want: "https://gitlab.com/test-group/test-sub-group/test-project/-/blob/commit/path/to/file#L1-2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileLineLink(tt.args.repositoryURL, tt.args.filename, tt.args.branch, tt.args.commit, tt.args.startLine, tt.args.endLine); got != tt.want {
				t.Errorf("GetFileLineLink() = %v, want %v", got, tt.want)
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
			name:         "GitLab main environment",
			envsFilePath: gitlabMainEnvsFilePath,
			want:         true,
		},
		{
			name:         "GitLab pr environment",
			envsFilePath: gitlabPrEnvsFilePath,
			want:         true,
		},
		{
			name:         "Not GitLab environment",
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
