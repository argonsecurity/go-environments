package utils

import (
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
