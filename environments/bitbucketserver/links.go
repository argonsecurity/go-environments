package bitbucketserver

import (
	"fmt"
)

func GetFileLink(repositoryURL string, filename string, branch string, commit string) string {
	if commit != "" {
		return fmt.Sprintf("%s/browse/%s?at=%s",
			repositoryURL,
			filename,
			commit)
	}
	if branch != "" {
		return fmt.Sprintf("%s/browse/%s?at=%s",
			repositoryURL,
			filename,
			branch)
	}

	return ""
}

func GetFileLineLink(repositoryURL string, filename string, branch string, commit string, startLine int, endLine int) string {
	url := GetFileLink(repositoryURL, filename, branch, commit)
	if startLine != 0 {
		lines := fmt.Sprintf("#%d", startLine)
		if endLine != 0 && endLine != startLine {
			lines = fmt.Sprintf("%s-%d", lines, endLine)
		}

		url += lines
	}

	return url
}
