package environments

import (
	"fmt"

	"github.com/argonsecurity/go-environments/enums"
	azure "github.com/argonsecurity/go-environments/environments/jenkins/environments/azure"
	azureserver "github.com/argonsecurity/go-environments/environments/jenkins/environments/azure_server"
	"github.com/argonsecurity/go-environments/environments/jenkins/environments/bitbucket"
	bitbucketserver "github.com/argonsecurity/go-environments/environments/jenkins/environments/bitbucket_server"
	"github.com/argonsecurity/go-environments/environments/jenkins/environments/gitlab"
	"github.com/argonsecurity/go-environments/models"
)

func EnhanceConfiguration(configuration *models.Configuration) *models.Configuration {
	if gitlab.IsCurrentEnvironment() || configuration.Repository.Source == enums.GitlabServer || configuration.Repository.Source == enums.Gitlab {
		return gitlab.EnhanceConfiguration(configuration)
	}

	if bitbucket.IsCurrentEnvironment() || configuration.Repository.Source == enums.Bitbucket {
		return bitbucket.EnhanceConfiguration(configuration)
	}

	if bitbucketserver.IsCurrentEnvironment() || configuration.Repository.Source == enums.BitbucketServer {
		return bitbucketserver.EnhanceConfiguration(configuration)
	}

	return configuration
}

func BuildGenericScmLink(baseUrl, org, subgroups, repo string, isSshUrl bool) string {
	return fmt.Sprintf("%s/%s/%s%s", baseUrl, org, subgroups, repo)
}

func BuildScmLink(baseUrl, org, subgroups, repo string, isSshUrl bool, repoSource enums.Source) string {
	if repoSource == enums.Azure {
		return azure.BuildScmLink(baseUrl, org, subgroups, repo, isSshUrl)
	}
	if repoSource == enums.AzureServer {
		return azureserver.BuildScmLink(baseUrl, org, subgroups, repo, isSshUrl)
	}
	if repoSource == enums.BitbucketServer {
		return bitbucketserver.BuildScmLink(baseUrl, org, subgroups, repo, isSshUrl)
	}

	return BuildGenericScmLink(baseUrl, org, subgroups, repo, isSshUrl)
}
