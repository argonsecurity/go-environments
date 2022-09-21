package bitbucketserver

import (
	"fmt"
	"os"
	"strings"

	"github.com/argonsecurity/go-utils/environments/models"
	"github.com/argonsecurity/go-utils/http"
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
