package global

import (
	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/azure"
	"github.com/argonsecurity/go-environments/environments/bitbucket"
	"github.com/argonsecurity/go-environments/environments/github"
	"github.com/argonsecurity/go-environments/environments/gitlab"
)

type GetFileLinkFunc func(string, string, string, string, int, int) string

func GetFileLink(source enums.Source, repositoryURL string, filename string, branch string, commit string, startLine int, endLine int) string {
	var f GetFileLinkFunc
	switch source {
	case enums.Github, enums.GithubServer:
		f = github.GetFileLink
	case enums.Gitlab, enums.GitlabServer:
		f = gitlab.GetFileLink
	case enums.Azure, enums.AzureServer:
		f = azure.GetFileLink
	case enums.Bitbucket, enums.BitbucketServer:
		f = bitbucket.GetFileLink
	}

	if f != nil {
		return f(repositoryURL, filename, branch, commit, startLine, endLine)
	}

	return ""
}
