package environments

import (
	"github.com/argonsecurity/go-environments/enums"
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
