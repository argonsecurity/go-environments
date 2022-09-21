package azureserver

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
			name: "Azure Server HTTP URL",
			args: args{
				baseUrl:   "https://azure-devops.aquaseclabs.com",
				org:       "DefaultCollection",
				subgroups: "argon-monitor/_git/",
				repo:      "billy-integration-tests",
				isSshUrl:  false,
			},
			want: "https://azure-devops.aquaseclabs.com/DefaultCollection/argon-monitor/_git/billy-integration-tests",
		},
		{
			name: "Azure Server SSH URL",
			args: args{
				baseUrl:   "https://azure-devops.aquaseclabs.com",
				org:       "DefaultCollection",
				subgroups: "argon-monitor/_git/",
				repo:      "billy-integration-tests",
				isSshUrl:  true,
			},
			want: "https://azure-devops.aquaseclabs.com/DefaultCollection/argon-monitor/_git/billy-integration-tests",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildScmLink(tt.args.baseUrl, tt.args.org, tt.args.subgroups, tt.args.repo, tt.args.isSshUrl)
			assert.Equal(t, tt.want, got)
		})
	}
}
