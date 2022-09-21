package bitbucket

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/argonsecurity/go-utils/environments/enums"
	"github.com/argonsecurity/go-utils/environments/environments/testutils"
	"github.com/argonsecurity/go-utils/environments/models"
	"github.com/stretchr/testify/assert"
)

var (
	bitbucketMainEnvsFilePath = "testdata/bitbucket-pipelines-main-env.json"
	bitbucketPrEnvsFilePath   = "testdata/bitbucket-pipelines-pr-env.json"
	testRepoPath              = "/tmp/bitbucket/repo"
	testRepoUrl               = "http://bitbucket.org/test-workspace/test-repo"
	testRepoCloneUrl          = fmt.Sprintf("%s%s", testRepoUrl, ".git")
	testdataPath              = "../bitbucket/testdata/repo"
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
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Bitbucket,
				},
				Organization: models.Entity{
					Name: "test-workspace",
				},
				Pipeline: models.Entity{
					Id:   "{052d6f7d-516e-4207-9e54-7446023ce285}",
					Name: "test-repo",
				},
				Run: models.Entity{
					Id: "2",
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
					Url:      testRepoUrl,
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Bitbucket,
				},
				Organization: models.Entity{
					Name: "test-workspace",
				},
				Pipeline: models.Entity{
					Id:   "{70045d3c-f44f-4507-8c6f-1f5e326a083a}",
					Name: "test-repo",
				},
				Run: models.Entity{
					Id: "5",
				},
				Runner: models.Runner{
					OS:           runtime.GOOS,
					Architecture: runtime.GOARCH,
				},
				PullRequest: models.PullRequest{
					Id: "3",
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

func Test_environment_GetFileLineLink(t *testing.T) {
	type args struct {
		filePath   string
		ref        string
		lineNumber int
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
				filePath:   "path/to/file",
				ref:        "branchName",
				lineNumber: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/branchName/path/to/file#lines-1",
		},
		{
			name: "File from branch with line number 0",
			args: args{
				filePath:   "path/to/file",
				ref:        "branchName",
				lineNumber: 0,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/branchName/path/to/file",
		},
		{
			name: "File from commit",
			args: args{
				filePath:   "path/to/file",
				ref:        "1a70bx6328bad78d919dca422d1as1g1ec97c5f6",
				lineNumber: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/1a70bx6328bad78d919dca422d1as1g1ec97c5f6/path/to/file#lines-1",
		},
		{
			name: "Empty file path",
			args: args{
				filePath:   "",
				ref:        "1a70bx6328bad78d919dca422d1as1g1ec97c5f6",
				lineNumber: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src/1a70bx6328bad78d919dca422d1as1g1ec97c5f6/#lines-1",
		},
		{
			name: "Empty ref",
			args: args{
				filePath:   "path/to/file",
				ref:        "",
				lineNumber: 1,
			},
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         "https://bitbucket.org/test-workspace/test-repo/src//path/to/file#lines-1",
		},
		{
			name: "Not Bitbucket environment",
			args: args{
				filePath:   "path/to/file",
				ref:        "branchName",
				lineNumber: 1,
			},
			envsFilePath: "",
			want:         "https://bitbucket.org//src/branchName/path/to/file#lines-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			if got := e.GetFileLineLink(tt.args.filePath, tt.args.ref, tt.args.lineNumber); got != tt.want {
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
			name:         "GitLab main environment",
			envsFilePath: bitbucketMainEnvsFilePath,
			want:         true,
		},
		{
			name:         "GitLab pr environment",
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
