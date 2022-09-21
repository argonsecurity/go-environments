package gitlab

import (
	"fmt"
	"os"
	"strings"

	"github.com/argonsecurity/go-environments/http"
	"github.com/argonsecurity/go-environments/models"
)

const (
	gitlabSourceRepoURLEnv          = "gitlabSourceRepoURL"
	gitlabTargetBranchEnv           = "gitlabTargetBranch"
	gitlabSourceRepoHttpUrlEnv      = "gitlabSourceRepoHttpUrl"
	gitlabUserUsernameEnv           = "gitlabUserUsername"
	gitlabMergeRequestLastCommitEnv = "gitlabMergeRequestLastCommit"
	gitlabBranchEnv                 = "gitlabBranch"
	gitlabSourceBranchEnv           = "gitlabSourceBranch"
	gitlabBeforeEnv                 = "gitlabBefore"
	gitlabSourceRepoNameEnv         = "gitlabSourceRepoName"
	gitlabSourceNamespaceEnv        = "gitlabSourceNamespace"
	gitlabUserNameEnv               = "gitlabUserName"
)

func EnhanceConfiguration(configuration *models.Configuration) *models.Configuration {
	if _, isExists := os.LookupEnv(gitlabSourceRepoHttpUrlEnv); isExists {
		return EnhanceConfigurationWithEnvs(configuration)
	}

	return configuration
}

func EnhanceConfigurationWithEnvs(configuration *models.Configuration) *models.Configuration {
	configuration.Repository.CloneUrl = os.Getenv(gitlabSourceRepoHttpUrlEnv)
	configuration.Repository.Name = os.Getenv(gitlabSourceRepoNameEnv)
	configuration.Branch = os.Getenv(gitlabBranchEnv)
	configuration.PullRequest.SourceRef.Branch = os.Getenv(gitlabSourceBranchEnv)
	configuration.PullRequest.TargetRef.Branch = os.Getenv(gitlabTargetBranchEnv)
	configuration.BeforeCommitSha = os.Getenv(gitlabBeforeEnv)
	configuration.CommitSha = os.Getenv(gitlabMergeRequestLastCommitEnv)
	configuration.Pusher.Username = os.Getenv(gitlabUserUsernameEnv)
	configuration.Pusher.Name = os.Getenv(gitlabUserNameEnv)

	return configuration
}

func IsCurrentEnvironment() bool {
	_, isExist := os.LookupEnv(gitlabSourceRepoURLEnv)
	return isExist
}

func CheckGitlabByHTTPRequest(url string, httpClient http.HTTPService) bool {
	apiUrl := fmt.Sprintf("%s/api/v4/users", url)
	_, err := httpClient.Get(apiUrl, nil, nil)
	return err != nil && strings.Contains(err.Error(), "403 Forbidden")
}
