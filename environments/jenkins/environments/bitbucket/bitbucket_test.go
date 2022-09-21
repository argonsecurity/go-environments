package bitbucket

import (
	"testing"

	"github.com/argonsecurity/go-environments/environments/testutils"
	"github.com/argonsecurity/go-environments/models"
	"github.com/stretchr/testify/assert"
)

const (
	jenkinsBitbucketPrPayloadEnvs = "testdata/jenkins-bitbucket-pr-payload-env.json"
	jenkinsBitbucketPrEnvs        = "testdata/jenkins-bitbucket-pr-env.json"
)

func TestEnhanceConfigurationWithPayload(t *testing.T) {
	type args struct {
		configuration *models.Configuration
	}
	tests := []struct {
		name         string
		envsFilePath string
		args         args
		want         *models.Configuration
	}{
		{
			name:         "Jenkins Bitbucket pr environment",
			envsFilePath: jenkinsBitbucketPrPayloadEnvs,
			args: args{
				configuration: &models.Configuration{},
			},
			want: &models.Configuration{
				Repository: models.Repository{
					Id:   "{60978281-75ff-46ec-b025-7069f4bb23ec}",
					Name: "test-repo",
					Url:  "https://bitbucket.org/test-workspace/test-repo",
				},
				PullRequest: models.PullRequest{
					Id: "4",
					SourceRef: models.Ref{
						Sha: "vhjuaz2qjoi3",
					},
					TargetRef: models.Ref{
						Sha: "8fw5f2zdmxf5",
					},
					Url: "https://bitbucket.org/test-workspace/test-repo/pull-requests/4",
				},
				Pusher: models.Pusher{
					Username: "User Name",
					Entity: models.Entity{
						Id: "{4f61e3a7-f8c3-43ae-a16c-d2ab56d72a8f}",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envCleanup := testutils.SetEnvsFromFile(tt.envsFilePath)
			t.Cleanup(envCleanup)
			got := EnhanceConfigurationWithPayload(tt.args.configuration)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEnhanceConfigurationWithEnvs(t *testing.T) {
	type args struct {
		configuration *models.Configuration
	}
	tests := []struct {
		name         string
		envsFilePath string
		args         args
		want         *models.Configuration
	}{
		{
			name:         "Jenkins Bitbucket pr environment",
			envsFilePath: jenkinsBitbucketPrEnvs,
			args: args{
				configuration: &models.Configuration{},
			},
			want: &models.Configuration{
				Branch: "test-branch",
				PullRequest: models.PullRequest{
					Id: "4",
					SourceRef: models.Ref{
						Branch: "test-branch",
					},
					TargetRef: models.Ref{
						Branch: "master",
					},
					Url: "https://bitbucket.org/test-workspace/test-repo/pull-requests/4",
				},
			},
		},
	}
	for _, tt := range tests {
		envCleanup := testutils.SetEnvsFromFile(tt.envsFilePath)
		t.Cleanup(envCleanup)
		t.Run(tt.name, func(t *testing.T) {
			got := EnhanceConfigurationWithEnvs(tt.args.configuration)
			assert.Equal(t, tt.want, got)
		})
	}
}
