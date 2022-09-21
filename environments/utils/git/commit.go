package git

func (gc *Client) GetGitCommit(repositoryPath string) (string, error) {
	return gc.GitExecInDir(repositoryPath, "rev-parse", "HEAD")

}
