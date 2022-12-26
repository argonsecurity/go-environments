package bitbucket

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/testutils"
	"github.com/argonsecurity/go-environments/models"
	"github.com/stretchr/testify/assert"
)

var (
	bitbucketMainEnvsFilePath = "testdata/bitbucket-pipelines-main-env.json"
	bitbucketPrEnvsFilePath   = "testdata/bitbucket-pipelines-pr-env.json"
	testRepoPath              = "/tmp/bitbucket/repo"
	testRepoUrl               = "http://bitbucket.org/test-workspace/test-repo"
	testRepoCloneUrl          = fmt.Sprintf("%s%s", testRepoUrl, ".git")
	testdataPath              = "../bitbucket/testdata/repo"

	testBranch          = "branch"
	testBranchWithSlash = "feature/branch"
	testCommit          = "commit"
	testPath            = "path/to/file"
)

func Test_environment_GetConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		envsFilePath string
		want         *models.Configuration
		wantErr      bool
	}{
		{
			name:         "Bitbucket main configuration",
			envsFilePath: bitbucketMainEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://bitbucket.org",
				SCMApiUrl: "https://api.bitbucket.org/2.0",
				LocalPath: testRepoPath,
				Branch:    "master",
				CommitSha: "vdnbxo7pmcoepieoyx82sxve4k9d7664joc9c6af",
				Repository: models.Repository{
					Id:       "{d41c6669-e5cb-4bfb-96f3-77bebd632437}",
					Name:     "test-repo",
					FullName: "test-workspace/test-repo",
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Bitbucket,
				},
				Organization: models.Entity{
					Name: "test-workspace",
				},
				Pipeline: models.Pipeline{
					Entity: models.Entity{
						Id:   "{052d6f7d-516e-4207-9e54-7446023ce285}",
						Name: "test-repo",
					},
					Path: "bitbucket-pipelines.yml",
				},
				Run: models.BuildRun{
					BuildId:     "2",
					BuildNumber: "2",
				},
				Runner: models.Runner{
					OS:           runtime.GOOS,
					Architecture: runtime.GOARCH,
				},
				PullRequest: models.PullRequest{
					Id: "",
				},
				PipelinePaths: []string{"/tmp/bitbucket/repo/bitbucket-pipelines.yml"},
				Environment:   enums.Bitbucket,
				ScmId:         "a664f15182cd78c6d563889694770ec9",
			},
			wantErr: false,
		},
		{
			name:         "Bitbucket pr configuration",
			envsFilePath: bitbucketPrEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://bitbucket.org",
				SCMApiUrl: "https://api.bitbucket.org/2.0",
				LocalPath: testRepoPath,
				Branch:    "test-branch",
				CommitSha: "44oo5siajopw",
				Repository: models.Repository{
					Id:       "{d41c6669-e5cb-4bfb-96f3-77bebd632437}",
					Name:     "test-repo",
					FullName: "test-workspace/test-repo",
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Bitbucket,
				},
				Organization: models.Entity{
					Name: "test-workspace",
				},
				Pipeline: models.Pipeline{
					Entity: models.Entity{
						Id:   "{70045d3c-f44f-4507-8c6f-1f5e326a083a}",
						Name: "test-repo",
					},
					Path: "bitbucket-pipelines.yml",
				},
				Run: models.BuildRun{
					BuildId:     "5",
					BuildNumber: "5",
				},
				Runner: models.Runner{
					OS:           runtime.GOOS,
					Architecture: runtime.GOARCH,
				},
				PullRequest: models.PullRequest{
					Id: "3",
					TargetRef: models.Ref{
						Sha:    "",
						Branch: "master",
					},
				},
				PipelinePaths: []string{"/tmp/bitbucket/repo/bitbucket-pipelines.yml"},
				Environment:   enums.Bitbucket,
				ScmId:         "a664f15182cd78c6d563889694770ec9",
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
			name:         "Bitbucket environment",
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/pipelines/results/2/steps/{5026c107-7d6c-4140-8824-7bfb990b3e78}",
		},
		{
			name:         "Not Bitbucket environment",
			envsFilePath: "",
			want:         "https://bitbucket.org//pipelines/results//steps/",
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
			name:         "Bitbucket environment",
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/pipelines/results/2",
		},
		{
			name:         "Not Bitbucket environment",
			envsFilePath: "",
			want:         "https://bitbucket.org//pipelines/results/",
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
		args         args
		envsFilePath string
		want         string
	}{
		{
			name: "File from branch",
			args: args{
				filename: testPath,
				branch:   testBranch,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/branch/path/to/file",
		},
		{
			name: "File from commit",
			args: args{
				filename: testPath,
				commit:   testCommit,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file",
		},
		{
			name: "Empty file path",
			args: args{
				filename: "",
				commit:   testCommit,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/commit/",
		},
		{
			name: "Empty ref",
			args: args{
				filename: testPath,
				branch:   "",
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src//path/to/file",
		},
		{
			name: "Not Bitbucket environment",
			args: args{
				filename: testPath,
				branch:   testBranch,
			},
			envsFilePath: "",
			want:         "https://bitbucket.org//src/branch/path/to/file",
		},
		{
			name: "Branch with Slash",
			args: args{
				filename: testPath,
				branch:   testBranchWithSlash,
				commit:   testCommit,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file?at=feature%2Fbranch",
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
			name: "File from branch",
			args: args{
				filePath:  testPath,
				branch:    testBranch,
				startLine: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/branch/path/to/file#lines-1",
		},
		{
			name: "File from branch with line number 0",
			args: args{
				filePath:  testPath,
				branch:    testBranch,
				startLine: 0,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/branch/path/to/file",
		},
		{
			name: "File from commit",
			args: args{
				filePath:  testPath,
				commit:    testCommit,
				startLine: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file#lines-1",
		},
		{
			name: "Empty file path",
			args: args{
				filePath:  "",
				commit:    testCommit,
				startLine: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/commit/#lines-1",
		},
		{
			name: "Empty ref",
			args: args{
				filePath:  testPath,
				branch:    "",
				startLine: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src//path/to/file#lines-1",
		},
		{
			name: "Not Bitbucket environment",
			args: args{
				filePath:  testPath,
				branch:    testBranch,
				startLine: 1,
			},
			envsFilePath: "",
			want:         "https://bitbucket.org//src/branch/path/to/file#lines-1",
		},
		{
			name: "Branch with Slash",
			args: args{
				filePath:  testPath,
				branch:    testBranchWithSlash,
				commit:    testCommit,
				startLine: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file?at=feature%2Fbranch#lines-1",
		},
		{
			name: "Same line",
			args: args{
				filePath:  testPath,
				commit:    testCommit,
				startLine: 1,
				endLine:   1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file#lines-1",
		},
		{
			name: "Different lines",
			args: args{
				filePath:  testPath,
				commit:    testCommit,
				startLine: 1,
				endLine:   2,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file#lines-1:2",
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
			want: "http://bitbucket.org/test-workspace/test-repo/src/branch/path/to/file",
		},
		{
			name: "With commit",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				commit:        testCommit,
			},
			want: "http://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file",
		},
		{
			name: "With commit and branch",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				commit:        testCommit,
				branch:        testBranch,
			},
			want: "http://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file?at=branch",
		},
		{
			name: "With commit and branch with slash",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				commit:        testCommit,
				branch:        testBranchWithSlash,
			},
			want: "http://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file?at=feature%2Fbranch",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileLink(tt.args.repositoryURL, tt.args.filename, tt.args.branch, tt.args.commit); got != tt.want {
				t.Errorf("GetFileLink() = %v, want %v", got, tt.want)
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
			name: "No line numbers",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				branch:        testBranch,
			},
			want: "http://bitbucket.org/test-workspace/test-repo/src/branch/path/to/file",
		},
		{
			name: "Same line",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				branch:        testBranch,
				startLine:     1,
				endLine:       1,
			},
			want: "http://bitbucket.org/test-workspace/test-repo/src/branch/path/to/file#lines-1",
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
			want: "http://bitbucket.org/test-workspace/test-repo/src/branch/path/to/file#lines-1:2",
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
			want: "http://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file#lines-1:2",
		},
		{
			name: "With commit and branch",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				commit:        testCommit,
				branch:        testBranch,
				startLine:     1,
				endLine:       2,
			},
			want: "http://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file?at=branch#lines-1:2",
		},
		{
			name: "With commit and branch with slash",
			args: args{
				repositoryURL: testRepoUrl,
				filename:      testPath,
				commit:        testCommit,
				branch:        testBranchWithSlash,
				startLine:     1,
				endLine:       2,
			},
			want: "http://bitbucket.org/test-workspace/test-repo/src/commit/path/to/file?at=feature%2Fbranch#lines-1:2",
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
			name:         "Bitbucket main environment",
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         true,
		},
		{
			name:         "Bitbucket pr environment",
			envsFilePath: bitbucketPrEnvsFilePath,
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
