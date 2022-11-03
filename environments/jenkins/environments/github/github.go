package github

import (
	"github.com/argonsecurity/go-environments/models"
	"os"
)

const (
	gitHubToken                      = "GITHUB_TOKEN"
	branchEnvGhSinglePipeline        = "ghprbSourceBranch"
	targetBranchNameGhSinglePipeline = "ghprbTargetBranch"
)

func IsCurrentEnvironment() bool {
	_, isExist := os.LookupEnv(gitHubToken)
	return isExist
}

func EnhanceConfiguration(configuration *models.Configuration) *models.Configuration {

	if configuration.PullRequest.TargetRef.Branch == "" {
		configuration.PullRequest.TargetRef.Branch = os.Getenv(targetBranchNameGhSinglePipeline)
	}

	if configuration.PullRequest.SourceRef.Branch == "" {
		configuration.PullRequest.SourceRef.Branch = os.Getenv(branchEnvGhSinglePipeline)
	}
	return configuration
}
