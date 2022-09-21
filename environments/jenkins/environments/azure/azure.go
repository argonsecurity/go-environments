package azure

import (
	"fmt"
	"strings"
)

func BuildScmLink(baseUrl, org, subgroups, repo string, isSshUrl bool) string {
	fixedBaseUrl := baseUrl
	fixedSubgroups := subgroups
	if isSshUrl {
		fixedBaseUrl = strings.Replace(baseUrl, "ssh.", "", 1)
		fixedSubgroups = fmt.Sprintf("%s_git/", subgroups)
	}
	return fmt.Sprintf("%s/%s/%s%s", fixedBaseUrl, org, fixedSubgroups, repo)
}
