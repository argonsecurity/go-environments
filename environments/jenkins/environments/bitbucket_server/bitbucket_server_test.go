package bitbucketserver

import (
	"testing"

	"github.com/argonsecurity/go-utils/environments/environments/testutils"
	"github.com/argonsecurity/go-utils/environments/models"
	"github.com/stretchr/testify/assert"
)

const jenkinsBitbucketServerPrEnvs = "testdata/jenkins-bitbucket-server-pr-env.json"

func TestEnhanceConfiguration(t *testing.T) {
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
			name:         "Jenkins Bitbucket Server pr env",
			envsFilePath: jenkinsBitbucketServerPrEnvs,
			args: args{
				configuration: &models.Configuration{},
			},
			want: &models.Configuration{
				PullRequest: models.PullRequest{
					Id:  "4",
					Url: "https://bitbucket.org/test-workspace/test-repo/pull-requests/4",
				},
				Pusher: models.Pusher{
					Username: "User Name",
					Email:    "user@email.com",
					Entity: models.Entity{
						Name: "username",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envCleanup := testutils.SetEnvsFromFile(tt.envsFilePath)
			t.Cleanup(envCleanup)
			got := EnhanceConfiguration(tt.args.configuration)
			assert.Equal(t, tt.want, got)
		})
	}
}
