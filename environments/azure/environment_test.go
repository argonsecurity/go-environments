package azure

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/argonsecurity/go-utils/environments/enums"
	"github.com/argonsecurity/go-utils/environments/environments/testutils"
	"github.com/argonsecurity/go-utils/environments/models"
	"github.com/stretchr/testify/assert"
)

var (
	azureMainEnvsFilePath = "testdata/azure-pipelines-main-env.json"
	azurePrEnvsFilePath   = "testdata/azure-pipelines-pr-env.json"
	testRepoPath          = "/tmp/azure/repo"
	testRepoUrl           = "http://dev.azure.com/test-workspace/test-repo"
	testRepoCloneUrl      = fmt.Sprintf("%s%s", testRepoUrl, ".git")
	testdataPath          = "../azure/testdata/repo"
)

func Test_environment_GetConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		envsFilePath string
		want         *models.Configuration
		wantErr      bool
	}{
		{
			name:         "Azure main environment",
			envsFilePath: azureMainEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://dev.azure.com/test-organization/",
				SCMApiUrl: "https://dev.azure.com/test-organization/",
				LocalPath: testRepoPath,
				Branch:    "refs/heads/main",
				ProjectId: "a65c82d2-643f-4362-b55d-ad095527b237",
				CommitSha: "7dutv00rz4u9ogrhcwzt7hcjn4rt2v9s6zoa065o",
				Organization: models.Entity{
					Id:   "45c695d9-7d50-483a-bf50-3cfcc8140293",
					Name: "test-organization",
				},
				Repository: models.Repository{
					Id:       "6613da8a-3e14-4d4e-a06b-f8933353e044",
					Name:     "test-repo",
					Url:      "https://test-organization@dev.azure.com/test-organization/test-repo/_git/test-repo",
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Azure,
				},
				Pusher: models.Pusher{
					Username: "User Name",
					Email:    "user@email.com",
				},
				Job: models.Entity{
					Id:   "Job",
					Name: "Job",
				},
				Pipeline: models.Entity{
					Id:   "a65c82d2-643f-4362-b55d-ad095527b237-8",
					Name: "test-pipeline",
				},
				Run: models.BuildRun{
					BuildId: "152",
					BuildNumber: "20220912.1",
				},
				Runner: models.Runner{
					Id:           "8",
					Name:         "Hosted Agent",
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
				PipelinePaths: []string{"/tmp/azure/repo/azure-pipelines.yml"},
				Environment:   enums.Azure,
				ScmId:         "7716833d1f05b3d746cfd34d72d0aa11",
			},
			wantErr: false,
		},
		{
			name:         "Azure pr environment",
			envsFilePath: azurePrEnvsFilePath,
			want: &models.Configuration{
				Url:       "https://dev.azure.com/test-organization/",
				SCMApiUrl: "https://dev.azure.com/test-organization/",
				LocalPath: testRepoPath,
				Branch:    "refs/heads/test-branch",
				ProjectId: "a65c82d2-643f-4362-b55d-ad095527b237",
				CommitSha: "1zu7szijr66vf093ih0b3rhj5tzl5tfs1mlih5yj",
				Organization: models.Entity{
					Id:   "45c695d9-7d50-483a-bf50-3cfcc8140293",
					Name: "test-organization",
				},
				Repository: models.Repository{
					Id:       "6613da8a-3e14-4d4e-a06b-f8933353e044",
					Name:     "test-repo",
					Url:      "https://test-organization@dev.azure.com/test-organization/test-repo/_git/test-repo",
					CloneUrl: testRepoCloneUrl,
					Source:   enums.Azure,
				},
				Pusher: models.Pusher{
					Username: "Microsoft.VisualStudio.Services.TFS",
					Email:    "user@email.com",
				},
				Job: models.Entity{
					Id:   "Job",
					Name: "Job",
				},
				Pipeline: models.Entity{
					Id:   "a65c82d2-643f-4362-b55d-ad095527b237-8",
					Name: "test-pipeline",
				},
				Run: models.BuildRun{
					BuildId:     "178",
					BuildNumber: "20220912.12",
				},
				Runner: models.Runner{
					Id:           "8",
					Name:         "Hosted Agent",
					OS:           "Linux",
					Distribution: "ubuntu20",
					Architecture: "X64",
				},
				PullRequest: models.PullRequest{
					Id: "37",
					SourceRef: models.Ref{
						Branch: "refs/heads/test-branch",
					},
					TargetRef: models.Ref{
						Branch: "refs/heads/main",
					},
				},
				PipelinePaths: []string{"/tmp/azure/repo/azure-pipelines.yml"},
				Environment:   enums.Azure,
				ScmId:         "7716833d1f05b3d746cfd34d72d0aa11",
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
			name:         "Azure environment",
			envsFilePath: azureMainEnvsFilePath,
			want:         "https://dev.azure.com/test-organization/test-repo/_build/results?buildId=152&view=logs&j=11c096ac-5eeb-4c25-bbd4-537e204d2a12&t=ea62e213-a757-44dc-aac1-09850e53c1d2",
		},
		{
			name:         "Not Azure environment",
			envsFilePath: "",
			want:         "/_build/results?buildId=&view=logs&j=&t=",
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
			name:         "Azure environment",
			envsFilePath: azureMainEnvsFilePath,
			want:         "https://dev.azure.com/test-organization/test-repo/_build?definitionId=8&_a=summary",
		},
		{
			name:         "Azure environment",
			envsFilePath: "",
			want:         "/_build?definitionId=&_a=summary",
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
		branchName string
		lineNumber int
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
				filePath:   "path/to/file",
				branchName: "branchName",
				lineNumber: 1,
			},
			envsFilePath: azureMainEnvsFilePath,
			want:         "https://dev.azure.com/test-organization/_git/test-repo?path=path%2Fto%2Ffile&version=GBbranchName&line=1&lineEnd=2&lineStartColumn=1&lineEndColumn=1&lineStyle=plain&_a=contents",
		},
		{
			name: "File from branch with line number 0",
			args: args{
				filePath:   "path/to/file",
				branchName: "branchName",
				lineNumber: 0,
			},
			envsFilePath: azureMainEnvsFilePath,
			want:         "https://dev.azure.com/test-organization/_git/test-repo?path=path%2Fto%2Ffile&version=GBbranchName&_a=contents",
		},
		{
			name: "File from commit",
			args: args{
				filePath:   "path/to/file",
				branchName: "1a70bx6328bad78d919dca422d1as1g1ec97c5f6",
				lineNumber: 1,
			},
			envsFilePath: azureMainEnvsFilePath,
			want:         "https://dev.azure.com/test-organization/_git/test-repo?path=path%2Fto%2Ffile&version=GB1a70bx6328bad78d919dca422d1as1g1ec97c5f6&line=1&lineEnd=2&lineStartColumn=1&lineEndColumn=1&lineStyle=plain&_a=contents",
		},
		{
			name: "Empty file path",
			args: args{
				filePath:   "",
				branchName: "1a70bx6328bad78d919dca422d1as1g1ec97c5f6",
				lineNumber: 1,
			},
			envsFilePath: azureMainEnvsFilePath,
			want:         "https://dev.azure.com/test-organization/_git/test-repo?path=&version=GB1a70bx6328bad78d919dca422d1as1g1ec97c5f6&line=1&lineEnd=2&lineStartColumn=1&lineEndColumn=1&lineStyle=plain&_a=contents",
		},
		{
			name: "Empty ref",
			args: args{
				filePath:   "path/to/file",
				branchName: "",
				lineNumber: 1,
			},
			envsFilePath: azureMainEnvsFilePath,
			want:         "https://dev.azure.com/test-organization/_git/test-repo?path=path%2Fto%2Ffile&version=GB&line=1&lineEnd=2&lineStartColumn=1&lineEndColumn=1&lineStyle=plain&_a=contents",
		},
		{
			name: "Not azure environment",
			args: args{
				filePath:   "path/to/file",
				branchName: "branchName",
				lineNumber: 1,
			},
			envsFilePath: "",
			want:         "_git/?path=path%2Fto%2Ffile&version=GBbranchName&line=1&lineEnd=2&lineStartColumn=1&lineEndColumn=1&lineStyle=plain&_a=contents",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := prepareTest(t, tt.envsFilePath)
			if got := e.GetFileLineLink(tt.args.filePath, tt.args.branchName, tt.args.lineNumber); got != tt.want {
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
			name:         "Azure main environment",
			envsFilePath: azureMainEnvsFilePath,
			want:         true,
		},
		{
			name:         "Azure pr environment",
			envsFilePath: azurePrEnvsFilePath,
			want:         true,
		},
		{
			name:         "Not Azure environment",
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

func Test_getOrganizationName(t *testing.T) {
	type args struct {
		collectionURI string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy path",
			args: args{
				collectionURI: "https://dev.azure.com/test-organization/",
			},
			want: "test-organization",
		},
		{
			name: "Empty collection URI",
			args: args{
				collectionURI: "",
			},
			want: "",
		},
		{
			name: "Invalid collection URI",
			args: args{
				collectionURI: "https://dev.azure.com/",
			},
			want: "",
		},
		{
			name: "Invalid URI",
			args: args{
				collectionURI: "invalid",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrganizationName(tt.args.collectionURI); got != tt.want {
				t.Errorf("getOrganizationName() = %v, want %v", got, tt.want)
			}
		})
	}
}
