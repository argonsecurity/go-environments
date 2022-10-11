package githubserver

import (
	"fmt"
	"github.com/argonsecurity/go-environments/http"
	"strings"
)

func CheckGithubServerByHTTPRequest(url string, httpClient http.HTTPService) bool {
	apiUrl := fmt.Sprintf("%s/api/v3/meta", url)
	_, err := httpClient.Get(apiUrl, nil, nil)
	return err == nil
}

func GetGithubServerApiUrl(url string) string {
	if strings.Contains(url, "/api/v3") {
		return strings.Trim(url, "/")
	}
	return fmt.Sprintf("%s/api/v3", strings.Trim(url, "/"))
}
