package bitbucketserver

import (
	"fmt"
	"os"
	"strings"

	"github.com/argonsecurity/go-environments/http"
	"github.com/argonsecurity/go-environments/models"
)

const (
	pusherNameEnv        = "CHANGE_AUTHOR"
	pusherEmailEnv       = "CHANGE_AUTHOR_EMAIL"
	pusherDisplayNameEnv = "CHANGE_AUTHOR_DISPLAY_NAME"
	pullRequestIdEnv     = "CHANGE_ID"
	pullRequestUrlEnv    = "CHANGE_URL"
)

func EnhanceConfiguration(configuration *models.Configuration) *models.Configuration {
	configuration.Pusher.Username = os.Getenv(pusherDisplayNameEnv)
	configuration.Pusher.Email = os.Getenv(pusherEmailEnv)
	configuration.Pusher.Entity.Name = os.Getenv(pusherNameEnv)
	configuration.PullRequest.Id = os.Getenv(pullRequestIdEnv)
	configuration.PullRequest.Url = os.Getenv(pullRequestUrlEnv)
	return configuration
}

func IsCurrentEnvironment() bool {
	_, isExist := os.LookupEnv("CHANGE_BRANCH")
	return isExist
}

func CheckBitbucketServerByHTTPRequest(url string, httpClient http.HTTPService) bool {
	apiUrl := fmt.Sprintf("%s/rest/api/1.0/users", url)
	_, err := httpClient.Get(apiUrl, nil, nil)
	return err == nil || strings.Contains(err.Error(), "bitbucket")
}

func BuildScmLink(baseUrl, org, subgroups, repo string, isSshUrl bool) string {
	fixedBaseUrl := baseUrl
	fixedSubgroups := fmt.Sprintf("%srepos/", subgroups)
	if isSshUrl {
		fixedBaseUrl = strings.Replace(fixedBaseUrl, "git@", "", 1)
	}
	fixedBaseUrl = fmt.Sprintf("%s/projects", fixedBaseUrl)
	return fmt.Sprintf("%s/%s/%s%s", fixedBaseUrl, org, fixedSubgroups, repo)
}
