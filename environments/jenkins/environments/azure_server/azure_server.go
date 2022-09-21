package azureserver

import (
	"fmt"
	"strings"
)

func BuildScmLink(baseUrl, org, subgroups, repo string, isSshUrl bool) string {
	fixedBaseUrl := baseUrl
	if isSshUrl {
		fixedBaseUrl = strings.Replace(baseUrl, "ssh.", "", 1)
	}
	return fmt.Sprintf("%s/%s/%s%s", fixedBaseUrl, org, subgroups, repo)
}
