package utils

import (
	"fmt"
	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/jenkins/environments/azure"
	azureserver "github.com/argonsecurity/go-environments/environments/jenkins/environments/azure_server"
	bitbucketserver "github.com/argonsecurity/go-environments/environments/jenkins/environments/bitbucket_server"
	"net/url"
	"regexp"
	"strings"
)

var (
	urlRegexp = regexp.MustCompile(`(https?:\/\/|git@)(.*?(?::\d+)?)[\/:](.+?)(?:\.git)`)

	bitbucketServerUriRegexp = regexp.MustCompile(`scm/(.*?)/(.*?)(?:\.git|$)`)
	uriRegexp                = regexp.MustCompile(`/?(.+?)/(?:(.+/))?(.+?)(?:\.git|$)`)
	httpUrlRegexp            = regexp.MustCompile(`(https?://.+?)(/.+)`)
	sshUrlRegexp             = regexp.MustCompile(`(ssh?://.+?)(?:\:[0-9]+)(/.+)`)
	gitUrlRegexp             = regexp.MustCompile(`git@(.+?)(\:.+)`)
	sshUriRegexp             = regexp.MustCompile(`(?:/(?:v3|[0-9]+))?/(?P<org>.+?)/(.+/)?(?P<repo>.+?)(?:\.git|$)`)
	sshIdentificationRegexp  = regexp.MustCompile(`^.*@|ssh://`)

	githubApiUrl    = "https://api.github.com"
	gitlabApiUrl    = "https://gitlab.com/api/v4"
	azureApiUrl     = ""
	bitbucketApiUrl = "https://api.bitbucket.org/2.0"
)

func ParseGitURL(gitUrl string) []string {
	result := urlRegexp.FindStringSubmatch(gitUrl)
	protocol := result[1]
	host := result[2]
	paths := result[3]

	if result[1] == "git@" {
		protocol = "https://"
	}

	urls := []string{fmt.Sprintf("%s%s", protocol, host)}
	for i, path := range strings.Split(paths, "/") {
		urls = append(urls, fmt.Sprintf("%s/%s", urls[i], path))
	}
	return urls
}

func StripCredentialsFromUrl(urlToStrip string) string {
	urlObject, err := url.Parse(urlToStrip)
	if err != nil {
		return urlToStrip
	}

	urlObject.User = nil
	return urlObject.String()
}

// ParseDataFromCloneUrl extracts data from the clone url
// and returns the repository url, organization, repository name and repository full name (including org, group, sub group etc.)
// the base url is used for cases where the base of the scm url includes a part of the URI
//
// i.e https://example.company.io/gitlab
func ParseDataFromCloneUrl(cloneUrl, apiUrl string, repoSource enums.Source) (string, string, string, string, error) {
	var regexp = uriRegexp
	baseUrl, uri, isSshUrl, err := getUriFromCloneUrl(cloneUrl, apiUrl)
	if err != nil {
		return "", "", "", "", err
	}

	// In bitbucket server the clone url looks like this: https://server-bitbucket.company.com/scm/project/repo.git
	// so we need to extract the organization and repository names using a different regex
	if isSshUrl {
		regexp = sshUriRegexp
	} else if repoSource == enums.BitbucketServer {
		regexp = bitbucketServerUriRegexp
	}
	results := regexp.FindAllStringSubmatch(uri, -1)
	if len(results) == 0 {
		return "", "", "", "", fmt.Errorf("could not parse clone url: %s", cloneUrl)
	}
	result := results[0]

	var org, subgroups, repo string
	if len(result) == 4 { // url contains subgroups
		org, subgroups, repo = result[1], result[2], result[3]
	} else { // url doesn't contains subgroups
		org, repo = result[1], result[2]
	}
	repositoryFullName := strings.TrimSuffix(strings.TrimPrefix(result[0], "/"), ".git")

	return buildScmLink(baseUrl, org, subgroups, repo, isSshUrl, repoSource), org, repo, repositoryFullName, nil
}

func buildGenericScmLink(baseUrl, org, subgroups, repo string, isSshUrl bool) string {
	return fmt.Sprintf("%s/%s/%s%s", baseUrl, org, subgroups, repo)
}

func buildScmLink(baseUrl, org, subgroups, repo string, isSshUrl bool, repoSource enums.Source) string {
	if repoSource == enums.Azure {
		return azure.BuildScmLink(baseUrl, org, subgroups, repo, isSshUrl)
	}
	if repoSource == enums.AzureServer {
		return azureserver.BuildScmLink(baseUrl, org, subgroups, repo, isSshUrl)
	}
	if repoSource == enums.BitbucketServer {
		return bitbucketserver.BuildScmLink(baseUrl, org, subgroups, repo, isSshUrl)
	}

	return buildGenericScmLink(baseUrl, org, subgroups, repo, isSshUrl)
}

// getUriFromCloneUrl for cases where the baseUrl is not actually
// a part of the cloneUrl (i.e. Github), we need to extract the URI
// from the cloneUrl without using the baseUrl
func getUriFromCloneUrl(cloneUrl, apiUrl string) (string, string, bool, error) {
	isSshUrl := sshIdentificationRegexp.MatchString(cloneUrl)
	if strings.Contains(cloneUrl, apiUrl) && apiUrl != "" {
		return apiUrl, strings.Replace(cloneUrl, apiUrl, "", 1), isSshUrl, nil
	}
	if httpUrlRegexp.MatchString(cloneUrl) {
		result := httpUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)
		return result[0][1], result[0][2], isSshUrl, nil
	}
	if sshUrlRegexp.MatchString(cloneUrl) {
		result := sshUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)
		return strings.Replace(result[0][1], "ssh", "https", 1), result[0][2], isSshUrl, nil
	}
	results := gitUrlRegexp.FindAllStringSubmatch(cloneUrl, -1)
	if len(results) == 0 {
		return "", "", isSshUrl, fmt.Errorf("could not parse clone url: %s", cloneUrl)
	}
	result := results[0]
	return fmt.Sprintf("https://%s", result[1]), strings.Replace(result[2], ":", "/", 1), isSshUrl, nil
}
