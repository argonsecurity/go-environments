package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var possibleUserEnvVars = []string{
	"BITBUCKET_ACTOR",
	"GITHUB_ACTOR",
	"CODEBUILD_GIT_AUTHOR",
	"CIRCLE_USERNAME",
}

func DetectPusher() string {
	for _, userEnv := range possibleUserEnvVars {
		if v, ok := os.LookupEnv(userEnv); ok {
			return v
		}
	}

	re := regexp.MustCompile(`(?m)^.*<(.+?)>`)
	path, _ := os.Getwd()
	logsHeadFile := filepath.Join(path, ".git", "logs", "HEAD")
	if _, err := os.Stat(logsHeadFile); err == nil {
		contents, err := os.ReadFile(logsHeadFile)
		if err == nil {
			matches := re.FindAllSubmatch(contents, -1)
			if len(matches) >= 1 {
				return string(matches[len(matches)-1][1])
			}
		}
	}

	if v, ok := os.LookupEnv("USERNAME"); ok {
		return fmt.Sprintf("Fallback: %s", v)
	}

	return ""
}
