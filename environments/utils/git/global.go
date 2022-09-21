package git

import (
	"fmt"
	"os"
	"regexp"
)

var (
	GlobalGitClient GitClient
)

func init() {
	var err error
	GlobalGitClient, err = InitClient("")
	if err != nil {
		panic(err)
	}
}

func GetGitRemoteURL(repositoryPath string) (string, error) {
	return GlobalGitClient.GetGitRemoteURL(repositoryPath)
}

func GetGitBranch(repositoryPath string, commit string) (string, error) {
	return GlobalGitClient.GetGitBranch(repositoryPath, commit)
}

func GetGitCommit(repositoryPath string) (string, error) {
	return GlobalGitClient.GetGitCommit(repositoryPath)
}

func AddRemoteUrl(repositoryPath string, remoteUrl string) error {
	return GlobalGitClient.AddRemoteUrl(repositoryPath, remoteUrl)
}

func CreateGitRepository(path string) error {
	return GlobalGitClient.CreateGitRepository(path)
}

func IsPathContainsRepository(path string) bool {
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		if _, err := os.Stat(fmt.Sprintf("%s/.git", path)); err == nil {
			return true
		}
	}
	return false
}

func TrimBranchName(branchName string) string {
	branchHeadRegexp := regexp.MustCompile(`^\w+/\w+/`)
	return branchHeadRegexp.ReplaceAllString(branchName, "")
}
