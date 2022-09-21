package git

import (
	"fmt"
	"strings"
)

func (gc *Client) GetGitBranch(repositoryPath string, commit string) (string, error) {
	branch, err := gc.GitExecInDir(repositoryPath, "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", err
	}
	if branch == "HEAD" { // this means we are running in detached mode
		return gc.getBranchContainingCommit(repositoryPath, commit)
	}

	return branch, nil
}

func (gc *Client) getBranchContainingCommit(repositoryPath, commit string) (string, error) {
	parsedOutput, err := gc.GitExecInDir(repositoryPath, "branch", "-a", "--contains", commit)
	if err != nil {
		return "", err
	}
	lines := strings.Split(parsedOutput, "\n")
	for _, line := range lines {

		if strings.Contains(line, "HEAD") { // skip detached reference
			continue
		}

		if strings.HasPrefix(line, "*") {
			return strings.TrimSpace(strings.TrimPrefix(line, "*")), nil
		}
		branch := strings.TrimSpace(line)
		headCommit, err := gc.getBranchHeadCommit(repositoryPath, branch)
		if err != nil {
			fmt.Printf("failed to get branch HEAD commit: %s %s", branch, err)
			continue
		}
		if headCommit == commit {
			return TrimBranchName(branch), nil
		}

	}
	return "", nil
}

func (gc *Client) getBranchHeadCommit(repositoryPath, branch string) (string, error) {
	return gc.GitExecInDir(repositoryPath, "rev-parse", branch)
}
