package jenkins

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/argonsecurity/go-utils/environments/enums"
	"github.com/argonsecurity/go-utils/environments/environments/testutils"
	"github.com/argonsecurity/go-utils/environments/environments/testutils/mocks"
	"github.com/argonsecurity/go-utils/environments/environments/utils/git"
	"github.com/argonsecurity/go-utils/environments/models"
	"github.com/stretchr/testify/assert"
)

var (
	jenkinsGithubMainFullEnvsFilePath    = "testdata/jenkins-github-main-full-env.json"
	jenkinsGithubMainMinimalEnvsFilePath = "testdata/jenkins-github-main-minimal-env.json"
	testRepoPath                         = "/tmp/jenkins/repo"
	testRepoUrl                          = "https://github.com/test-organization/test-repo"
	testRepoCloneUrl                     = fmt.Sprintf("%s%s", testRepoUrl, ".git")
	testRepoCommit                       = "kcy8v2oazy4acuo1475rgtzg0nh403p23vw812lv"
	testdataPath                         = "../jenkins/testdata/repo"
)

func Test_parseDataFromCloneUrl(t *testing.T) {
	type args struct {
		cloneUrl   string
		apiUrl     string
		repoSource enums.Source
	}
	tests := []struct {
		name     string
		args     args
		wantUrl  string
		wantOrg  string
		wantRepo string
	}{
		{
			name: "GitHub HTTP clone url",
			args: args{
				cloneUrl:   "https://github.com/test-organization/test-repo.git",
				apiUrl:     githubApiUrl,
				repoSource: enums.Github,
			},
			wantUrl:  "https://github.com/test-organization/test-repo",
			wantOrg:  "test-organization",
			wantRepo: "test-repo",
		},
		{
			name: "GitHub SSH clone url",
			args: args{
				cloneUrl:   "git@github.com:test-organization/test-repo.git",
				apiUrl:     githubApiUrl,
				repoSource: enums.Github,
			},
			wantUrl:  "https://github.com/test-organization/test-repo",
			wantOrg:  "test-organization",
			wantRepo: "test-repo",
		},
		{
			name: "GitLab HTTP clone url",
			args: args{
				cloneUrl:   "https://gitlab.com/test-group/subgroup1/subgroup2/test-project.git",
				apiUrl:     gitlabApiUrl,
				repoSource: enums.Gitlab,
			},
			wantUrl:  "https://gitlab.com/test-group/subgroup1/subgroup2/test-project",
			wantOrg:  "test-group",
			wantRepo: "test-project",
		},
		{
			name: "GitLab SSH clone url",
			args: args{
				cloneUrl:   "git@gitlab.com:test-group/subgroup1/subgroup2/test-project.git",
				apiUrl:     gitlabApiUrl,
				repoSource: enums.Gitlab,
			},
			wantUrl:  "https://gitlab.com/test-group/subgroup1/subgroup2/test-project",
			wantOrg:  "test-group",
			wantRepo: "test-project",
		},
		{
			name: "GitLab Server HTTP clone url",
			args: args{
				cloneUrl:   "https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project.git",
				apiUrl:     "https://server.com/gitlab",
				repoSource: enums.GitlabServer,
			},
			wantUrl:  "https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project",
			wantOrg:  "test-group",
			wantRepo: "test-project",
		},
		// {
		// 	name: "GitLab Server SSH clone url",
		// 	args: args{
		// 		cloneUrl:   "git@server.com:gitlab/test-group/subgroup1/subgroup2/test-project.git",
		// 		apiUrl:     gitlabApiUrl,
		// 		repoSource: enums.Gitlab,
		// 	},
		// 	wantUrl:  "https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project",
		// 	wantOrg:  "test-group",
		// 	wantRepo: "test-project",
		// },
		{
			name: "Azure HTTP clone url",
			args: args{
				cloneUrl:   "https://dev.azure.com/test-organization/test-project/_git/test-repo",
				apiUrl:     azureApiUrl,
				repoSource: enums.Azure,
			},
			wantUrl:  "https://dev.azure.com/test-organization/test-project/_git/test-repo",
			wantOrg:  "test-organization",
			wantRepo: "test-repo",
		},
		// {
		// 	name: "Azure SSH clone url",
		// 	args: args{
		// 		cloneUrl:   "git@ssh.dev.azure.com:v3/test-organization/test-project/test-repo",
		// 		apiUrl:     azureApiUrl,
		// 		repoSource: enums.Azure,
		// 	},
		// 	wantUrl:  "https://dev.azure.com/test-organization/test-project/_git/test-repo",
		// 	wantOrg:  "test-organization",
		// 	wantRepo: "test-repo",
		// },
		// {
		// 	name: "Azure Server HTTP clone url",
		// 	args: args{
		// 		cloneUrl:   "https://dev.azure.com/test-organization/test-project/_git/test-repo",
		// 		apiUrl:     "",
		// 		repoSource: enums.AzureServer,
		// 	},
		// 	wantUrl:  "https://dev.azure.com/test-organization/test-project/_git/test-repo",
		// 	wantOrg:  "test-organization",
		// 	wantRepo: "test-repo",
		// },
		// {
		// 	name: "Azure Server SSH clone url",
		// 	args: args{
		// 		cloneUrl:   "git@ssh.dev.azure.com:v3/test-organization/test-project/test-repo",
		// 		apiUrl:     "",
		// 		repoSource: enums.AzureServer,
		// 	},
		// 	wantUrl:  "https://dev.azure.com/test-organization/test-project/_git/test-repo",
		// 	wantOrg:  "test-organization",
		// 	wantRepo: "test-repo",
		// },
		{
			name: "Bitbucket HTTP clone url",
			args: args{
				cloneUrl:   "https://bitbucket.org/test-project/test-repo.git",
				apiUrl:     bitbucketApiUrl,
				repoSource: enums.Bitbucket,
			},
			wantUrl:  "https://bitbucket.org/test-project/test-repo",
			wantOrg:  "test-project",
			wantRepo: "test-repo",
		},
		{
			name: "Bitbucket SSH clone url",
			args: args{
				cloneUrl:   "git@bitbucket.org:test-project/test-repo.git",
				apiUrl:     bitbucketApiUrl,
				repoSource: enums.Bitbucket,
			},
			wantUrl:  "https://bitbucket.org/test-project/test-repo",
			wantOrg:  "test-project",
			wantRepo: "test-repo",
		},
		{
			name: "Bitbucket Server HTTP clone url",
			args: args{
				cloneUrl:   "https://bitbucket.server.com/scm/TS/test-repo.git",
				apiUrl:     "https://bitbucket.server.com",
				repoSource: enums.BitbucketServer,
			},
			wantUrl:  "https://bitbucket.server.com/TS/test-repo",
			wantOrg:  "TS",
			wantRepo: "test-repo",
		},
		// {
		// 	name: "Bitbucket Server SSH clone url",
		// 	args: args{
		// 		cloneUrl:   "ssh://git@bitbucket.server.com:7999/TS/test-repo.git",
		// 		apiUrl:     "https://bitbucket.server.com",
		// 		repoSource: enums.BitbucketServer,
		// 	},
		// 	wantUrl:  "https://bitbucket.org/test-project/test-repo",
		// 	wantOrg:  "TS",
		// 	wantRepo: "test-repo",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUrl, gotOrg, gotRepo := parseDataFromCloneUrl(tt.args.cloneUrl, tt.args.apiUrl, tt.args.repoSource)
			if gotUrl != tt.wantUrl {
				t.Errorf("parseDataFromCloneUrl() gotUrl = %v, want %v", gotUrl, tt.wantUrl)
			}
			if gotOrg != tt.wantOrg {
				t.Errorf("parseDataFromCloneUrl() gotOrg = %v, want %v", gotOrg, tt.wantOrg)
			}
			if gotRepo != tt.wantRepo {
				t.Errorf("parseDataFromCloneUrl() gotRepo = %v, want %v", gotRepo, tt.wantRepo)
			}
		})
	}
}

func Test_getUriFromCloneUrl(t *testing.T) {
	type args struct {
		cloneUrl string
		apiUrl   string
	}
	tests := []struct {
		name        string
		args        args
		wantBaseUrl string
		wantUri     string
	}{
		{
			name: "CloneUrl contains apiUrl",
			args: args{
				cloneUrl: "https://server.com/test/org/repo.git",
				apiUrl:   "https://server.com/test",
			},
			wantBaseUrl: "https://server.com/test",
			wantUri:     "/org/repo.git",
		},
		{
			name: "CloneUrl does't contain apiUrl HTTP",
			args: args{
				cloneUrl: "https://github.com/org/repo.git",
				apiUrl:   "https://api.github.com",
			},
			wantBaseUrl: "https://github.com",
			wantUri:     "/org/repo.git",
		},
		{
			name: "CloneUrl does't contain apiUrl SSH",
			args: args{
				cloneUrl: "git@github.com:org/repo.git",
				apiUrl:   "https://api.github.com",
			},
			wantBaseUrl: "https://github.com",
			wantUri:     "/org/repo.git",
		},
		{
			name: "CloneUrl with subgroups does't contain apiUrl HTTP",
			args: args{
				cloneUrl: "https://gitlab.com/group/subgroup/repo.git",
				apiUrl:   "https://gitlab.com.com/api/v4",
			},
			wantBaseUrl: "https://gitlab.com",
			wantUri:     "/group/subgroup/repo.git",
		},
		{
			name: "CloneUrl with subgroups does't contain apiUrl SSH",
			args: args{
				cloneUrl: "git@gitlab.com:group/subgroup/repo.git",
				apiUrl:   "https://gitlab.com/api/v4",
			},
			wantBaseUrl: "https://gitlab.com",
			wantUri:     "/group/subgroup/repo.git",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBaseUrl, gotUri := getUriFromCloneUrl(tt.args.cloneUrl, tt.args.apiUrl)
			if gotBaseUrl != tt.wantBaseUrl {
				t.Errorf("getUriFromCloneUrl() gotBaseUrl = %v, want %v", gotBaseUrl, tt.wantBaseUrl)
			}
			if gotUri != tt.wantUri {
				t.Errorf("getUriFromCloneUrl() gotUri = %v, want %v", gotUri, tt.wantUri)
			}
		})
	}
}

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
				Pipeline: models.Entity{
					Id:   "test-project",
					Name: "test-project",
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
				Pipeline: models.Entity{
					Id:   "test-project",
					Name: "test-project",
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
		filename string
		ref      string
		line     int
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
			if got := e.GetFileLineLink(tt.args.filename, tt.args.ref, tt.args.line); got != tt.want {
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
