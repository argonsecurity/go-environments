package bitbucketserver

import (
	"testing"
)

const (
	testRepoURL  = "https://bitbucket-server.com/projects/ar/repos/reponame"
	testFilename = "path/to/file"
	testBranch   = "branch"
	testCommit   = "commit"
)

func TestGetFileLink(t *testing.T) {
	type args struct {
		repositoryURL string
		filename      string
		branch        string
		commit        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "With branch",
			args: args{
				repositoryURL: testRepoURL,
				filename:      testFilename,
				branch:        testBranch,
			},
			want: "https://bitbucket-server.com/projects/ar/repos/reponame/browse/path/to/file?at=branch",
		},
		{
			name: "With commit",
			args: args{
				repositoryURL: testRepoURL,
				filename:      testFilename,
				commit:        testCommit,
			},
			want: "https://bitbucket-server.com/projects/ar/repos/reponame/browse/path/to/file?at=commit",
		},
		{
			name: "With branch and commit",
			args: args{
				repositoryURL: testRepoURL,
				filename:      testFilename,
				branch:        testBranch,
				commit:        testCommit,
			},
			want: "https://bitbucket-server.com/projects/ar/repos/reponame/browse/path/to/file?at=commit",
		},
		{
			name: "Empty repo url",
			args: args{
				repositoryURL: "",
				filename:      testFilename,
				branch:        testBranch,
			},
			want: "/browse/path/to/file?at=branch",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileLink(tt.args.repositoryURL, tt.args.filename, tt.args.branch, tt.args.commit); got != tt.want {
				t.Errorf("GetFileLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileLineLink(t *testing.T) {
	type args struct {
		repositoryURL string
		filename      string
		branch        string
		commit        string
		startLine     int
		endLine       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No lines",
			args: args{
				repositoryURL: testRepoURL,
				filename:      testFilename,
				branch:        testBranch,
			},
			want: "https://bitbucket-server.com/projects/ar/repos/reponame/browse/path/to/file?at=branch",
		},
		{
			name: "Only start line",
			args: args{
				repositoryURL: testRepoURL,
				filename:      testFilename,
				commit:        testCommit,
				startLine:     1,
			},
			want: "https://bitbucket-server.com/projects/ar/repos/reponame/browse/path/to/file?at=commit#1",
		},
		{
			name: "With start line and end line",
			args: args{
				repositoryURL: testRepoURL,
				filename:      testFilename,
				branch:        testBranch,
				commit:        testCommit,
				startLine:     1,
				endLine:       2,
			},
			want: "https://bitbucket-server.com/projects/ar/repos/reponame/browse/path/to/file?at=commit#1-2",
		},
		{
			name: "Same start and end line",
			args: args{
				repositoryURL: testRepoURL,
				filename:      testFilename,
				branch:        testBranch,
				startLine:     1,
				endLine:       1,
			},
			want: "https://bitbucket-server.com/projects/ar/repos/reponame/browse/path/to/file?at=branch#1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileLineLink(tt.args.repositoryURL, tt.args.filename, tt.args.branch, tt.args.commit, tt.args.startLine, tt.args.endLine); got != tt.want {
				t.Errorf("GetFileLineLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
