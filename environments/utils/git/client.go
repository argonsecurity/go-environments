package git

import "os/exec"

type GitClient interface {
	GetGitRemoteURL(repositoryPath string) (string, error)
	AddRemoteUrl(repositoryPath string, remoteUrl string) error
	GetGitCommit(repositoryPath string) (string, error)
	GetGitBranch(repositoryPath string, commit string) (string, error)
	CreateGitRepository(path string) error
}

type Client struct {
	binPath string
}

func InitClient(gitPath string) (*Client, error) {
	var err error
	if gitPath == "" {
		gitPath, err = exec.LookPath("git")
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		binPath: gitPath,
	}, nil
}
