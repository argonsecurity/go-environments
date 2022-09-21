package environments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			got := BuildGenericScmLink(tt.args.baseUrl, tt.args.org, tt.args.subgroups, tt.args.repo, tt.args.isSshUrl)
			assert.Equal(t, tt.want, got)
		})
	}
}
