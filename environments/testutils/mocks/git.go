package mocks

import "github.com/argonsecurity/go-utils/environments/environments/utils/git"

type MockGitClient struct {
	remoteUrl string
	commit    string
	branch    string

	commandResult string

	err error
}

// Setters
func (m *MockGitClient) SetRemoteUrl(remoteUrl string) *MockGitClient {
	m.remoteUrl = remoteUrl
	return m
}

func (m *MockGitClient) SetCommit(commit string) *MockGitClient {
	m.commit = commit
	return m
}

func (m *MockGitClient) SetBranch(branch string) *MockGitClient {
	m.branch = branch
	return m
}

func (m *MockGitClient) SetError(err error) *MockGitClient {
	m.err = err
	return m
}

func (m *MockGitClient) SetCommandResult(result string) *MockGitClient {
	m.commandResult = result
	return m
}

// Implementations
func (m *MockGitClient) GetGitRemoteURL(repositoryPath string) (string, error) {
	return m.remoteUrl, m.err
}

func (m *MockGitClient) GetGitCommit(repositoryPath string) (string, error) {
	return m.commit, m.err
}

func (m *MockGitClient) GetGitBranch(repositoryPath string, commit string) (string, error) {
	return m.branch, m.err
}

func (m *MockGitClient) AddRemoteUrl(repositoryPath string, remoteUrl string) error {
	return m.err
}
func (m *MockGitClient) CreateGitRepository(path string) error {
	return m.err
}

func (m *MockGitClient) GitExec(args ...string) (string, error) {
	return m.commandResult, m.err
}

func (m *MockGitClient) GitExecInDir(dir string, args ...string) (string, error) {
	return m.commandResult, m.err
}

// Set the global git client to the mock
func SetGitMock(mock *MockGitClient) {
	git.GlobalGitClient = mock
}

// Make sure MockGitClient implements GitClient
var _ git.GitClient = &MockGitClient{}
