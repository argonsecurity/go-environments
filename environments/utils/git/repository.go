package git

func (gc *Client) CreateGitRepository(path string) error {
	_, err := gc.GitExec("init", "--initial-branch=main", path)
	return err
}
