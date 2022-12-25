package utils

import (
	"github.com/argonsecurity/go-environments/enums"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGitURL(t *testing.T) {
	type args struct {
		gitUrl string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "GitHub HTTP clone url",
			args: args{
				gitUrl: "https://github.com/test-organization/test-repo.git",
			},
			want: []string{
				"https://github.com",
				"https://github.com/test-organization",
				"https://github.com/test-organization/test-repo",
			},
		},
		{
			name: "GitHub SSH clone url",
			args: args{
				gitUrl: "git@github.com:test-organization/test-repo.git",
			},
			want: []string{
				"https://github.com",
				"https://github.com/test-organization",
				"https://github.com/test-organization/test-repo",
			},
		},
		{
			name: "GitLab HTTP clone url",
			args: args{
				gitUrl: "https://gitlab.com/test-group/subgroup1/subgroup2/test-project.git",
			},
			want: []string{
				"https://gitlab.com",
				"https://gitlab.com/test-group",
				"https://gitlab.com/test-group/subgroup1",
				"https://gitlab.com/test-group/subgroup1/subgroup2",
				"https://gitlab.com/test-group/subgroup1/subgroup2/test-project",
			},
		},
		{
			name: "GitLab SSH clone url",
			args: args{
				gitUrl: "git@gitlab.com:test-group/subgroup1/subgroup2/test-project.git",
			},
			want: []string{
				"https://gitlab.com",
				"https://gitlab.com/test-group",
				"https://gitlab.com/test-group/subgroup1",
				"https://gitlab.com/test-group/subgroup1/subgroup2",
				"https://gitlab.com/test-group/subgroup1/subgroup2/test-project",
			},
		},
		{
			name: "GitLab Server HTTP clone url",
			args: args{
				gitUrl: "https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project.git",
			},
			want: []string{
				"https://server.com",
				"https://server.com/gitlab",
				"https://server.com/gitlab/test-group",
				"https://server.com/gitlab/test-group/subgroup1",
				"https://server.com/gitlab/test-group/subgroup1/subgroup2",
				"https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project",
			},
		},
		{
			name: "GitLab Server SSH clone url",
			args: args{
				gitUrl: "git@server.com:gitlab/test-group/subgroup1/subgroup2/test-project.git",
			},
			want: []string{
				"https://server.com",
				"https://server.com/gitlab",
				"https://server.com/gitlab/test-group",
				"https://server.com/gitlab/test-group/subgroup1",
				"https://server.com/gitlab/test-group/subgroup1/subgroup2",
				"https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project",
			},
		},
		// {
		// 	name: "Azure HTTP clone url",
		// 	args: args{
		// 		gitUrl: "https://dev.azure.com/test-organization/test-project/_git/test-repo",
		// 	},
		// 	want: []string{
		// 		"https://dev.azure.com",
		// 		"https://dev.azure.com/test-organization",
		// 		"https://dev.azure.com/test-organization/test-project",
		// 		"https://dev.azure.com/test-organization/test-project/_git",
		// 		"https://dev.azure.com/test-organization/test-project/_git/test-repo",
		// 	},
		// },
		// {
		// 	name: "Azure SSH clone url",
		// 	args: args{
		// 		gitUrl:   "git@ssh.dev.azure.com:v3/test-organization/test-project/test-repo",
		//
		//
		// 	},
		// 	want: []string{},
		//
		//
		// },
		// {
		// 	name: "Azure Server HTTP clone url",
		// 	args: args{
		// 		gitUrl:   "https://dev.azure.com/test-organization/test-project/_git/test-repo",
		//
		//
		// 	},
		// 	want: []string{},
		//
		//
		// },
		// {
		// 	name: "Azure Server SSH clone url",
		// 	args: args{
		// 		gitUrl:   "git@ssh.dev.azure.com:v3/test-organization/test-project/test-repo",
		//
		//
		// 	},
		// 	want: []string{},
		//
		//
		// },
		{
			name: "Bitbucket HTTP clone url",
			args: args{
				gitUrl: "https://bitbucket.org/test-project/test-repo.git",
			},
			want: []string{
				"https://bitbucket.org",
				"https://bitbucket.org/test-project",
				"https://bitbucket.org/test-project/test-repo",
			},
		},
		{
			name: "Bitbucket SSH clone url",
			args: args{
				gitUrl: "git@bitbucket.org:test-project/test-repo.git",
			},
			want: []string{
				"https://bitbucket.org",
				"https://bitbucket.org/test-project",
				"https://bitbucket.org/test-project/test-repo",
			},
		},
		{
			name: "Bitbucket Server HTTP clone url",
			args: args{
				gitUrl: "https://bitbucket.server.com/scm/TS/test-repo.git",
			},
			want: []string{
				"https://bitbucket.server.com",
				"https://bitbucket.server.com/scm",
				"https://bitbucket.server.com/scm/TS",
				"https://bitbucket.server.com/scm/TS/test-repo",
			},
		},
		// {
		// 	name: "Bitbucket Server SSH clone url",
		// 	args: args{
		// 		gitUrl:   "ssh://git@bitbucket.server.com:7999/TS/test-repo.git",
		//
		//
		// 	},
		// 	want: []string{},
		//
		//
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseGitURL(tt.args.gitUrl)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func TestStripCredentialsFromUrl(t *testing.T) {
	type args struct {
		urlToStrip string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Invalid url",
			args: args{
				urlToStrip: "https:user:pass@test.com/test1/test2",
			},
			want: "https:user:pass@test.com/test1/test2",
		},
		{
			name: "HTTP Url with no credentials",
			args: args{
				urlToStrip: "https://test.com/test1/test2",
			},
			want: "https://test.com/test1/test2",
		},
		{
			name: "HTTP Url with username only",
			args: args{
				urlToStrip: "https://user@test.com/test1/test2",
			},
			want: "https://test.com/test1/test2",
		},
		{
			name: "HTTP Url with password only",
			args: args{
				urlToStrip: "https://:pass@test.com/test1/test2",
			},
			want: "https://test.com/test1/test2",
		},
		{
			name: "HTTP Url with username and password only",
			args: args{
				urlToStrip: "https://user:pass@test.com/test1/test2",
			},
			want: "https://test.com/test1/test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StripCredentialsFromUrl(tt.args.urlToStrip)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBuildScmLink(t *testing.T) {
	type args struct {
		baseUrl   string
		org       string
		subgroups string
		repo      string
		isSshUrl  bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Github HTTP / SSH URL",
			args: args{
				baseUrl:   "https://github.com",
				org:       "argonsecurity",
				subgroups: "",
				repo:      "billy-integration-tests",
				isSshUrl:  false,
			},
			want: "https://github.com/argonsecurity/billy-integration-tests",
		},
		{
			name: "Gitlab HTTP / SSH URL",
			args: args{
				baseUrl:   "https://gitlab.com",
				org:       "dev-argon",
				subgroups: "",
				repo:      "billy-integration-tests",
				isSshUrl:  false,
			},
			want: "https://gitlab.com/dev-argon/billy-integration-tests",
		},
		{
			name: "Bitbucket HTTP / SSH URL",
			args: args{
				baseUrl:   "https://bitbucket.org",
				org:       "test-build",
				subgroups: "",
				repo:      "billy-integration-tests",
				isSshUrl:  false,
			},
			want: "https://bitbucket.org/test-build/billy-integration-tests",
		},
		{
			name: "Gitlab server HTTP / SSH URL",
			args: args{
				baseUrl:   "https://gitlab.aquasec.com",
				org:       "argon-monitor",
				subgroups: "",
				repo:      "billy-integration-tests",
				isSshUrl:  false,
			},
			want: "https://gitlab.aquasec.com/argon-monitor/billy-integration-tests",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildGenericScmLink(tt.args.baseUrl, tt.args.org, tt.args.subgroups, tt.args.repo, tt.args.isSshUrl)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseDataFromCloneUrl(t *testing.T) {
	type args struct {
		cloneUrl   string
		apiUrl     string
		repoSource enums.Source
	}
	tests := []struct {
		name             string
		args             args
		wantUrl          string
		wantOrg          string
		wantRepo         string
		wantRepoFullName string
		wantErr          bool
	}{
		{
			name: "GitHub HTTP clone url",
			args: args{
				cloneUrl:   "https://github.com/test-organization/test-repo.git",
				apiUrl:     githubApiUrl,
				repoSource: enums.Github,
			},
			wantUrl:          "https://github.com/test-organization/test-repo",
			wantOrg:          "test-organization",
			wantRepo:         "test-repo",
			wantRepoFullName: "test-organization/test-repo",
		},
		{
			name: "GitHub SSH clone url",
			args: args{
				cloneUrl:   "git@github.com:test-organization/test-repo.git",
				apiUrl:     githubApiUrl,
				repoSource: enums.Github,
			},
			wantUrl:          "https://github.com/test-organization/test-repo",
			wantOrg:          "test-organization",
			wantRepo:         "test-repo",
			wantRepoFullName: "test-organization/test-repo",
		},
		{
			name: "GitLab HTTP clone url",
			args: args{
				cloneUrl:   "https://gitlab.com/test-group/subgroup1/subgroup2/test-project.git",
				apiUrl:     gitlabApiUrl,
				repoSource: enums.Gitlab,
			},
			wantUrl:          "https://gitlab.com/test-group/subgroup1/subgroup2/test-project",
			wantOrg:          "test-group",
			wantRepo:         "test-project",
			wantRepoFullName: "test-group/subgroup1/subgroup2/test-project",
		},
		{
			name: "GitLab SSH clone url",
			args: args{
				cloneUrl:   "git@gitlab.com:test-group/subgroup1/subgroup2/test-project.git",
				apiUrl:     gitlabApiUrl,
				repoSource: enums.Gitlab,
			},
			wantUrl:          "https://gitlab.com/test-group/subgroup1/subgroup2/test-project",
			wantOrg:          "test-group",
			wantRepo:         "test-project",
			wantRepoFullName: "test-group/subgroup1/subgroup2/test-project",
		},
		{
			name: "GitLab Server HTTP clone url",
			args: args{
				cloneUrl:   "https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project.git",
				apiUrl:     "https://server.com/gitlab",
				repoSource: enums.GitlabServer,
			},
			wantUrl:          "https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project",
			wantOrg:          "test-group",
			wantRepo:         "test-project",
			wantRepoFullName: "test-group/subgroup1/subgroup2/test-project",
		},
		{
			name: "GitLab Server SSH clone url",
			args: args{
				cloneUrl:   "git@server.com/gitlab:test-group/subgroup1/subgroup2/test-project.git",
				apiUrl:     "https://server.com/gitlab",
				repoSource: enums.GitlabServer,
			},
			wantUrl:          "https://server.com/gitlab/test-group/subgroup1/subgroup2/test-project",
			wantOrg:          "test-group",
			wantRepo:         "test-project",
			wantRepoFullName: "test-group/subgroup1/subgroup2/test-project",
		},
		{
			name: "Azure HTTP clone url",
			args: args{
				cloneUrl:   "https://dev.azure.com/test-organization/test-project/_git/test-repo",
				apiUrl:     azureApiUrl,
				repoSource: enums.Azure,
			},
			wantUrl:          "https://dev.azure.com/test-organization/test-project/_git/test-repo",
			wantOrg:          "test-organization",
			wantRepo:         "test-repo",
			wantRepoFullName: "test-organization/test-project/_git/test-repo",
		},
		{
			name: "Azure SSH clone url",
			args: args{
				cloneUrl:   "git@ssh.dev.azure.com:v3/test-organization/test-project/test-repo",
				apiUrl:     azureApiUrl,
				repoSource: enums.Azure,
			},
			wantUrl:          "https://dev.azure.com/test-organization/test-project/_git/test-repo",
			wantOrg:          "test-organization",
			wantRepo:         "test-repo",
			wantRepoFullName: "v3/test-organization/test-project/test-repo",
		},
		{
			name: "Azure Server HTTP clone url",
			args: args{
				cloneUrl:   "https://azure-devops.server.com/test-organization/test-project/_git/test-repo",
				apiUrl:     "",
				repoSource: enums.AzureServer,
			},
			wantUrl:          "https://azure-devops.server.com/test-organization/test-project/_git/test-repo",
			wantOrg:          "test-organization",
			wantRepo:         "test-repo",
			wantRepoFullName: "test-organization/test-project/_git/test-repo",
		},
		{
			name: "Azure Server SSH clone url",
			args: args{
				cloneUrl:   "ssh://azure-devops.server.com:22/test-organization/test-project/_git/test-repo",
				apiUrl:     "",
				repoSource: enums.AzureServer,
			},
			wantUrl:          "https://azure-devops.server.com/test-organization/test-project/_git/test-repo",
			wantOrg:          "test-organization",
			wantRepo:         "test-repo",
			wantRepoFullName: "test-organization/test-project/_git/test-repo",
		},
		{
			name: "Bitbucket HTTP clone url",
			args: args{
				cloneUrl:   "https://bitbucket.org/test-project/test-repo.git",
				apiUrl:     bitbucketApiUrl,
				repoSource: enums.Bitbucket,
			},
			wantUrl:          "https://bitbucket.org/test-project/test-repo",
			wantOrg:          "test-project",
			wantRepo:         "test-repo",
			wantRepoFullName: "test-project/test-repo",
		},
		{
			name: "Bitbucket SSH clone url",
			args: args{
				cloneUrl:   "git@bitbucket.org:test-project/test-repo.git",
				apiUrl:     bitbucketApiUrl,
				repoSource: enums.Bitbucket,
			},
			wantUrl:          "https://bitbucket.org/test-project/test-repo",
			wantOrg:          "test-project",
			wantRepo:         "test-repo",
			wantRepoFullName: "test-project/test-repo",
		},
		{
			name: "Bitbucket Server HTTP clone url",
			args: args{
				cloneUrl:   "https://bitbucket.server.com/scm/TS/test-repo.git",
				apiUrl:     "https://bitbucket.server.com",
				repoSource: enums.BitbucketServer,
			},
			wantUrl:          "https://bitbucket.server.com/projects/TS/repos/test-repo",
			wantOrg:          "TS",
			wantRepo:         "test-repo",
			wantRepoFullName: "scm/TS/test-repo",
		},
		{
			name: "Bitbucket Server SSH clone url",
			args: args{
				cloneUrl:   "ssh://git@bitbucket.server.com:7999/TS/test-repo.git",
				apiUrl:     "https://bitbucket.server.com",
				repoSource: enums.BitbucketServer,
			},
			wantUrl:          "https://bitbucket.server.com/projects/TS/repos/test-repo",
			wantOrg:          "TS",
			wantRepo:         "test-repo",
			wantRepoFullName: "TS/test-repo",
		},
		{
			name: "Cannot parse clone url",
			args: args{
				cloneUrl:   "hello",
				apiUrl:     "",
				repoSource: enums.BitbucketServer,
			},
			wantUrl:          "",
			wantOrg:          "",
			wantRepo:         "",
			wantRepoFullName: "",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUrl, gotOrg, gotRepo, gotRepoFullName, err := ParseDataFromCloneUrl(tt.args.cloneUrl, tt.args.apiUrl, tt.args.repoSource)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDataFromCloneUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotUrl != tt.wantUrl {
				t.Errorf("ParseDataFromCloneUrl() gotUrl = %v, want %v", gotUrl, tt.wantUrl)
			}
			if gotOrg != tt.wantOrg {
				t.Errorf("ParseDataFromCloneUrl() gotOrg = %v, want %v", gotOrg, tt.wantOrg)
			}
			if gotRepo != tt.wantRepo {
				t.Errorf("ParseDataFromCloneUrl() gotRepo = %v, want %v", gotRepo, tt.wantRepo)
			}
			if gotRepoFullName != tt.wantRepoFullName {
				t.Errorf("ParseDataFromCloneUrl() gotRepoFullName = %v, want %v", gotRepoFullName, tt.wantRepoFullName)
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
		wantErr     bool
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
		{
			name: "Cannot parse clone url",
			args: args{
				cloneUrl: "hello",
				apiUrl:   "",
			},
			wantBaseUrl: "",
			wantUri:     "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBaseUrl, gotUri, _, err := getUriFromCloneUrl(tt.args.cloneUrl, tt.args.apiUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUriFromCloneUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotBaseUrl != tt.wantBaseUrl {
				t.Errorf("getUriFromCloneUrl() gotBaseUrl = %v, want %v", gotBaseUrl, tt.wantBaseUrl)
			}
			if gotUri != tt.wantUri {
				t.Errorf("getUriFromCloneUrl() gotUri = %v, want %v", gotUri, tt.wantUri)
			}
		})
	}
}
