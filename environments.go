package environments

import (
	"fmt"
	"github.com/argonsecurity/go-environments/environments/circleci"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/azure"
	"github.com/argonsecurity/go-environments/environments/bitbucket"
	"github.com/argonsecurity/go-environments/environments/bitbucketserver"
	"github.com/argonsecurity/go-environments/environments/github"
	"github.com/argonsecurity/go-environments/environments/gitlab"
	"github.com/argonsecurity/go-environments/environments/jenkins"
	"github.com/argonsecurity/go-environments/environments/localhost"
	"github.com/argonsecurity/go-environments/models"
)

var (
	environmentMapping map[enums.Source]Environment = map[enums.Source]Environment{
		enums.Github:    github.Github,
		enums.Gitlab:    gitlab.Gitlab,
		enums.Azure:     azure.Azure,
		enums.Bitbucket: bitbucket.Bitbucket,
		enums.Jenkins:   jenkins.Jenkins,
		enums.CircleCi:  circleci.CircleCi,
		enums.Localhost: localhost.Localhost,
	}
)

type GetFileLineLinkFunc func(string, string, string, string, int, int) string
type GetFileLinkFunc func(string, string, string, string) string

// Environment is an interface for interacting with CI/CD environments
type Environment interface {
	// GetConfiguration get a environment configuration
	GetConfiguration() (*models.Configuration, error)

	// GetBuildLink get a link to the current build
	GetBuildLink() string

	// GetStepLink get a link to the current step
	GetStepLink() string

	// GetFileLink get a link to a file
	GetFileLink(filename string, ref string, commit string) string

	// GetFileLineLink get a link to a file line
	GetFileLineLink(filename string, ref string, commit string, startLine int, endLine int) string

	// Name get the name of the environment
	Name() string

	// IsCurrentEnvironment detects if the runtime environment matches the object
	IsCurrentEnvironment() bool
}

// GetEnvironment get environment object that matches the name
func GetEnvironment(name string) (Environment, error) {
	if env, ok := environmentMapping[enums.Source(name)]; ok {
		return env, nil
	}
	return nil, fmt.Errorf("environment %s does not exist", name)
}

// DetectEnvironment get environment by detecting
func DetectEnvironment() Environment {
	for _, env := range environmentMapping {
		if env.IsCurrentEnvironment() {
			return env
		}
	}
	return localhost.Localhost
}

func GetOrDetectEnvironment(name string) (Environment, error) {
	if name != "" {
		return GetEnvironment(name)
	}
	return DetectEnvironment(), nil
}

func GetFileLineLink(source enums.Source, repositoryURL string, filename string, branch string, commit string, startLine int, endLine int) string {
	var f GetFileLineLinkFunc
	switch source {
	case enums.Github, enums.GithubServer:
		f = github.GetFileLineLink
	case enums.Gitlab, enums.GitlabServer:
		f = gitlab.GetFileLineLink
	case enums.Azure, enums.AzureServer:
		f = azure.GetFileLineLink
	case enums.Bitbucket:
		f = bitbucket.GetFileLineLink
	case enums.BitbucketServer:
		f = bitbucketserver.GetFileLineLink
	}

	if f != nil {
		return f(repositoryURL, filename, branch, commit, startLine, endLine)
	}

	return ""
}

func GetFileLink(source enums.Source, repositoryURL string, filename string, branch string, commit string) string {
	var f GetFileLinkFunc
	switch source {
	case enums.Github, enums.GithubServer:
		f = github.GetFileLink
	case enums.Gitlab, enums.GitlabServer:
		f = gitlab.GetFileLink
	case enums.Azure, enums.AzureServer:
		f = azure.GetFileLink
	case enums.Bitbucket:
		f = bitbucket.GetFileLink
	case enums.BitbucketServer:
		f = bitbucketserver.GetFileLink
	}

	if f != nil {
		return f(repositoryURL, filename, branch, commit)
	}

	return ""
}
