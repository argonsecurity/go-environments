package git

import (
	"errors"
	"strings"
)

const (
	remoteNameIndex = 0
	remoteURLIndex  = 1
)

func (gc *Client) GetGitRemoteURL(repositoryPath string) (string, error) {
	remotes, err := gc.getGitRemotes(repositoryPath)
	if err != nil {
		return "", err
	}

	if len(remotes) == 0 {
		return "", errors.New("no git remotes found")
	}

	for _, remote := range remotes {
		if remote[remoteNameIndex] == "origin" {
			return remote[remoteURLIndex], nil
		}
	}

	return remotes[0][remoteURLIndex], nil
}

func (gc *Client) AddRemoteUrl(repositoryPath string, remoteUrl string) error {
	_, err := gc.GitExecInDir(repositoryPath, "remote", "add", "origin", remoteUrl)
	return err
}

// getGitRemotes lists the git remotes in the given repository
// in the current format: ["REMOTE_NAME", "REMOTE_URL", "REMOTE_ACTION"]
func (gc *Client) getGitRemotes(repositoryPath string) ([][]string, error) {
	parsedOutput, err := gc.GitExecInDir(repositoryPath, "remote", "-v")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(parsedOutput, "\n")
	remotes := [][]string{}
	for _, line := range lines {
		remotes = append(remotes, strings.Fields(line))
	}
	return remotes, nil
}
