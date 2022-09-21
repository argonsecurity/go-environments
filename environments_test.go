package environments

import (
	"testing"

	"github.com/argonsecurity/go-utils/environments/environments/azure"
	"github.com/argonsecurity/go-utils/environments/environments/bitbucket"
	"github.com/argonsecurity/go-utils/environments/environments/github"
	"github.com/argonsecurity/go-utils/environments/environments/gitlab"
	"github.com/argonsecurity/go-utils/environments/environments/jenkins"
	"github.com/argonsecurity/go-utils/environments/environments/localhost"
	"github.com/argonsecurity/go-utils/environments/environments/testutils"
	"github.com/stretchr/testify/assert"
)

func TestGetEnvironment(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Environment
		wantErr bool
	}{
		{
			name: "GitHub environment",
			args: args{
				name: "github",
			},
			want:    github.Github,
			wantErr: false,
		},
		{
			name: "GitLab environment",
			args: args{
				name: "gitlab",
			},
			want:    gitlab.Gitlab,
			wantErr: false,
		},
		{
			name: "Azure environment",
			args: args{
				name: "azure",
			},
			want:    azure.Azure,
			wantErr: false,
		},
		{
			name: "Bitbucket environment",
			args: args{
				name: "bitbucket",
			},
			want:    bitbucket.Bitbucket,
			wantErr: false,
		},
		{
			name: "Jenkins environment",
			args: args{
				name: "jenkins",
			},
			want:    jenkins.Jenkins,
			wantErr: false,
		},
		{
			name: "Wrong environment",
			args: args{
				name: "test",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEnvironment(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDetectEnvironment(t *testing.T) {
	tests := []struct {
		name         string
		envsFilePath string
		want         Environment
	}{
		{
			name:         "GitHub environment",
			envsFilePath: "environments/github/testdata/github-workflows-main-env.json",
			want:         github.Github,
		},
		{
			name:         "GitLab environment",
			envsFilePath: "environments/gitlab/testdata/gitlab-ci-main-env.json",
			want:         gitlab.Gitlab,
		},
		{
			name:         "Azure environment",
			envsFilePath: "environments/azure/testdata/azure-pipelines-main-env.json",
			want:         azure.Azure,
		},
		{
			name:         "Bitbucket environment",
			envsFilePath: "environments/bitbucket/testdata/bitbucket-pipelines-main-env.json",
			want:         bitbucket.Bitbucket,
		},
		{
			name:         "Jenkins environment",
			envsFilePath: "environments/jenkins/testdata/jenkins-github-main-full-env.json",
			want:         jenkins.Jenkins,
		},
		{
			name:         "Other environment",
			envsFilePath: "",
			want:         localhost.Localhost,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envCleanup := testutils.SetEnvsFromFile(tt.envsFilePath)
			t.Cleanup(envCleanup)
			got := DetectEnvironment()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetOrDetectEnvironment(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name         string
		args         args
		envsFilePath string
		want         Environment
		wantErr      bool
	}{
		{
			name: "Get existing env",
			args: args{
				name: "github",
			},
			want:    github.Github,
			wantErr: false,
		},
		{
			name: "Get non existing env",
			args: args{
				name: "test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Detect existing env",
			args: args{
				name: "",
			},
			envsFilePath: "environments/github/testdata/github-workflows-main-env.json",
			want:         github.Github,
			wantErr:      false,
		},
		{
			name: "Detect non existing env",
			args: args{
				name: "",
			},
			want:    localhost.Localhost,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envCleanup := testutils.SetEnvsFromFile(tt.envsFilePath)
			t.Cleanup(envCleanup)
			got, err := GetOrDetectEnvironment(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrDetectEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
