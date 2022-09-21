package bitbucket

import (
	"encoding/json"
	"os"

	"github.com/argonsecurity/go-environments/environments/utils"
	"github.com/argonsecurity/go-environments/models"
)

const (
	bitbucketPayloadEnv         = "BITBUCKET_PAYLOAD"
	bitbucketSourceBranchEnv    = "CHANGE_BRANCH"
	bitbucketTargetBranchEnv    = "CHANGE_TARGET"
	bitbucketPullRequestLinkEnv = "CHANGE_URL"
	bitbucketPullRequestIdEnv   = "CHANGE_ID"
)

func EnhanceConfiguration(configuration *models.Configuration) *models.Configuration {
	if _, isExists := os.LookupEnv(bitbucketPayloadEnv); isExists {
		configuration = EnhanceConfigurationWithPayload(configuration)
	}

	return EnhanceConfigurationWithEnvs(configuration)
}

func EnhanceConfigurationWithPayload(configuration *models.Configuration) *models.Configuration {
	payload, err := initPayload()

	if err != nil {
		return configuration
	}

	configuration.PullRequest.Id = payload.PullRequest.Id
	configuration.PullRequest.Url = payload.PullRequest.Links.Html.Href
	configuration.PullRequest.SourceRef.Sha = payload.PullRequest.Source.Commit.Hash
	configuration.PullRequest.TargetRef.Sha = payload.PullRequest.Destination.Commit.Hash
	configuration.Pusher.Username = payload.PullRequest.Author.DisplayName
	configuration.Pusher.Entity.Id = payload.PullRequest.Author.Uuid
	configuration.Repository.Id = payload.Repository.Uuid
	configuration.Repository.Name = payload.Repository.Name
	configuration.Repository.Url = payload.Repository.Links.Html.Href

	return configuration
}

func EnhanceConfigurationWithEnvs(configuration *models.Configuration) *models.Configuration {
	utils.SetValueIfEnvExist(configuration, bitbucketSourceBranchEnv, "Branch")
	utils.SetValueIfEnvExist(configuration, bitbucketSourceBranchEnv, "PullRequest.SourceRef.Branch")
	utils.SetValueIfEnvExist(configuration, bitbucketTargetBranchEnv, "PullRequest.TargetRef.Branch")
	utils.SetValueIfEnvExist(configuration, bitbucketPullRequestLinkEnv, "PullRequest.Url")
	utils.SetValueIfEnvExist(configuration, bitbucketPullRequestIdEnv, "PullRequest.Id")

	return configuration
}

func initPayload() (*BitbucketPayload, error) {
	var payload *BitbucketPayload

	data, exists := os.LookupEnv(bitbucketPayloadEnv)

	if !exists {
		return nil, nil
	}

	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func IsCurrentEnvironment() bool {
	_, isExist := os.LookupEnv(bitbucketPayloadEnv)
	return isExist
}
