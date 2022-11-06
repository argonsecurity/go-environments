package github

import (
	"testing"

	"github.com/argonsecurity/go-environments/environments/testutils"
	"github.com/argonsecurity/go-environments/models"
	"github.com/stretchr/testify/assert"
)

const (
	jenkinsGitHubPrEnvs = "testdata/jenkins-github-pr-env.json"
)

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
			name:         "Jenkins GitHub pr environment",
			envsFilePath: jenkinsGitHubPrEnvs,
			args: args{
				configuration: &models.Configuration{},
			},
			want: &models.Configuration{
				PullRequest: models.PullRequest{
					SourceRef: models.Ref{
						Branch: "baseBranchTest",
					},
					TargetRef: models.Ref{
						Branch: "targetBranchTest",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envCleanup := testutils.SetEnvsFromFile(tt.envsFilePath)
			t.Cleanup(envCleanup)
			got := EnhanceConfigurationSinglePipeline(tt.args.configuration)
			assert.Equal(t, tt.want, got)
		})
	}
}
