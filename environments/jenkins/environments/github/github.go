package github

import (
	"github.com/argonsecurity/go-environments/environments/utils"
	"github.com/argonsecurity/go-environments/models"
	"os"
)

const (
	branchEnvGhSinglePipeline        = "ghprbSourceBranch"
	targetBranchNameGhSinglePipeline = "ghprbTargetBranch"
)

func IsGitHubSingleJenkinsPipeline() bool {
	_, isExist := os.LookupEnv(branchEnvGhSinglePipeline)
	return isExist
}

func EnhanceConfigurationSinglePipeline(configuration *models.Configuration) *models.Configuration {

	utils.SetValueIfEnvExist(configuration, branchEnvGhSinglePipeline, "PullRequest.SourceRef.Branch")
	utils.SetValueIfEnvExist(configuration, targetBranchNameGhSinglePipeline, "PullRequest.TargetRef.Branch")

	return configuration
}
