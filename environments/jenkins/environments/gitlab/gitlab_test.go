package gitlab

import (
	"testing"

	"github.com/argonsecurity/go-utils/environments/environments/testutils"
	"github.com/argonsecurity/go-utils/environments/models"
	"github.com/stretchr/testify/assert"
)

const (
	jenkinsGitlabMainEnvs = "testdata/jenkins-gitlab-main-env.json"
	jenkinsGitlabPrEnvs   = "testdata/jenkins-gitlab-pr-env.json"
)

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
			name:         "Jenkins GitLab main environment",
			envsFilePath: jenkinsGitlabMainEnvs,
			args: args{
				configuration: &models.Configuration{},
			},
			want: &models.Configuration{
				Repository: models.Repository{
					Name:     "test-project",
					CloneUrl: "https://gitlab.com/test-group/test-subgroup/test-project.git",
				},
				Branch: "main",
				PullRequest: models.PullRequest{
					SourceRef: models.Ref{
						Branch: "main",
					},
					TargetRef: models.Ref{
						Branch: "main",
					},
				},
				BeforeCommitSha: "69heyghp2item2hn18zw0cspsevl48ija3qlw8im",
				CommitSha:       "vkj76f5q3qxgxa7ssclxdiepx94pied8x6c72mjj",
				Pusher: models.Pusher{
					Username: "username1",
					Entity: models.Entity{
						Name: "User Name",
					},
				},
			},
		},
		{
			name:         "Jenkins GitLab pr environment",
			envsFilePath: jenkinsGitlabPrEnvs,
			args: args{
				configuration: &models.Configuration{},
			},
			want: &models.Configuration{
				Repository: models.Repository{
					Name:     "test-project",
					CloneUrl: "https://gitlab.com/test-group/test-subgroup/test-project.git",
				},
				Branch: "test-branch",
				PullRequest: models.PullRequest{
					SourceRef: models.Ref{
						Branch: "test-branch",
					},
					TargetRef: models.Ref{
						Branch: "main",
					},
				},
				CommitSha: "ms8c5gtsja31yyymwy1cfenxfvww90bbidcxvs3r",
				Pusher: models.Pusher{
					Entity: models.Entity{
						Name: "User Name",
					},
				},
			},
		},
		{
			name:         "Not Jenkins GitLab environment",
			envsFilePath: "",
			args: args{
				configuration: &models.Configuration{},
			},
			want: &models.Configuration{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envCleanup := testutils.SetEnvsFromFile(tt.envsFilePath)
			t.Cleanup(envCleanup)
			got := EnhanceConfigurationWithEnvs(tt.args.configuration)
			assert.Equal(t, tt.want, got)
		})
	}
}
