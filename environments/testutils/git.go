package testutils

import (
	"fmt"
	"os"

	"github.com/argonsecurity/go-environments/environments/utils/git"
	"github.com/otiai10/copy"
)

func PrepareTestGitRepository(repositoryPath string, repositoryUrl, testdataPath string) (cleanup func()) {
	os.RemoveAll(repositoryPath)
	if err := git.CreateGitRepository(repositoryPath); err != nil {
		panic(err)
	}

	if err := git.AddRemoteUrl(repositoryPath, repositoryUrl); err != nil {
		panic(err)
	}

	if err := copy.Copy(testdataPath, repositoryPath); err != nil {
		panic(err)
	}

	return func() {
		if err := os.RemoveAll(repositoryPath); err != nil {
			fmt.Printf("Failed to clean test repo %s - %s", repositoryPath, err)
		}
	}
}
