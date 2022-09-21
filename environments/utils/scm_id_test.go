package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateScmId(t *testing.T) {
	type args struct {
		cloneUrl string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GitHub HTTP clone url",
			args: args{
				cloneUrl: "https://github.com/argonsecurity/argon-utils.git",
			},
			want: "f2e46a756099ea7774015283dbe1a3de",
		},
		{
			name: "GitHub SSH clone url",
			args: args{
				cloneUrl: "git@github.com:argonsecurity/argon-utils.git",
			},
			want: "f2e46a756099ea7774015283dbe1a3de",
		},
		{
			name: "Gitlab HTTP clone url",
			args: args{
				cloneUrl: "https://gitlab.com/dev-argon/billy-integration-tests.git",
			},
			want: "45a9eabf9d10338117566ea40a3b0c00",
		},
		{
			name: "Gitlab SSH clone url",
			args: args{
				cloneUrl: "git@gitlab.com:dev-argon/billy-integration-tests.git",
			},
			want: "45a9eabf9d10338117566ea40a3b0c00",
		},
		{
			name: "Bitbucket HTTP clone url",
			args: args{
				cloneUrl: "https://bitbucket.org/test-build/billy-integration-tests.git",
			},
			want: "71335cd18a839126be1f85c97e56a5cd",
		},
		{
			name: "Bitbucket SSH clone url",
			args: args{
				cloneUrl: "git@bitbucket.org:test-build/billy-integration-tests.git",
			},
			want: "71335cd18a839126be1f85c97e56a5cd",
		},
		{
			name: "Azure HTTP clone url",
			args: args{
				cloneUrl: "https://argon-monitor@dev.azure.com/argon-monitor/billy-integration-tests/_git/billy-integration-tests",
			},
			want: "c305af8b0af714242fee8d24522657a6",
		},
		{
			name: "Azure SSH clone url",
			args: args{
				cloneUrl: "git@ssh.dev.azure.com:v3/argon-monitor/billy-integration-tests/billy-integration-tests",
			},
			want: "075896c7478716f6fa0e472733f70cf7", // ToDo: fix this to be the same as the HTTP clone url (ssh. && v3) https://scalock.atlassian.net/browse/SAAS-9107
		},
		{
			name: "Gitlab server HTTP clone url",
			args: args{
				cloneUrl: "https://gitlab.aquasec.com/argon-monitor/billy-integration-tests.git",
			},
			want: "27dea414bfe24d64706a012a014c3ee0",
		},
		{
			name: "Gitlab server SSH clone url",
			args: args{
				cloneUrl: "git@gitlab.aquasec.com:argon-monitor/billy-integration-tests.git",
			},
			want: "27dea414bfe24d64706a012a014c3ee0",
		},
		{
			name: "Bitbucket server HTTP clone url",
			args: args{
				cloneUrl: "https://bitbucket5.aquaseclabs.com/scm/ar/billy-integration-tests.git",
			},
			want: "adfb2526b5809e658faeb14ef943eeb4",
		},
		{
			name: "Bitbucket server SSH clone url",
			args: args{
				cloneUrl: "ssh://git@bitbucket5.aquaseclabs.com:7999/ar/billy-integration-tests.git",
			},
			want: "0540406957207529743eaa88615d1b81", // ToDo: fix this to be the same as the HTTP clone url (remove port and scm) https://scalock.atlassian.net/browse/SAAS-9107
		},
		{
			name: "Azure server HTTP clone url",
			args: args{
				cloneUrl: "http://azure-devops.aquaseclabs.com/DefaultCollection/argon-monitor/_git/billy-integration-tests",
			},
			want: "306e6b49b391148cd3ce28b315ed6d47",
		},
		{
			name: "Azure server SSH clone url",
			args: args{
				cloneUrl: "ssh://azure-devops.aquaseclabs.com:22/DefaultCollection/argon-monitor/_git/billy-integration-tests",
			},
			want: "18ae4e68ea47ef77340055d112511a5d", // ToDo: fix this to be the same as the HTTP clone url (remove port) https://scalock.atlassian.net/browse/SAAS-9107
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateScmId(tt.args.cloneUrl)
			assert.Equal(t, tt.want, got)
		})
	}
}
