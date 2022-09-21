package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	urlRegexp = regexp.MustCompile(`(https?:\/\/|git@)(.*?(?::\d+)?)[\/:](.+?)(?:\.git)`)
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
