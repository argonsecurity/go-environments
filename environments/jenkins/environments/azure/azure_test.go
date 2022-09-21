package azure

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
			name: "Azure HTTP URL",
			args: args{
				baseUrl:   "https://dev.azure.com",
				org:       "argon-monitor",
				subgroups: "billy-integration-tests/_git/",
				repo:      "billy-integration-tests",
				isSshUrl:  false,
			},
			want: "https://dev.azure.com/argon-monitor/billy-integration-tests/_git/billy-integration-tests",
		},
		{
			name: "Azure SSH URL",
			args: args{
				baseUrl:   "https://ssh.dev.azure.com",
				org:       "argon-monitor",
				subgroups: "billy-integration-tests/",
				repo:      "billy-integration-tests",
				isSshUrl:  true,
			},
			want: "https://dev.azure.com/argon-monitor/billy-integration-tests/_git/billy-integration-tests",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildScmLink(tt.args.baseUrl, tt.args.org, tt.args.subgroups, tt.args.repo, tt.args.isSshUrl)
			assert.Equal(t, tt.want, got)
		})
	}
}
