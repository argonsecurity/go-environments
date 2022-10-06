package localhost

import (
	"os"

	"github.com/argonsecurity/go-environments/enums"
	"github.com/argonsecurity/go-environments/environments/utils"
	"github.com/argonsecurity/go-environments/environments/utils/git"
	"github.com/argonsecurity/go-environments/models"
)

var (
	Localhost     = environment{}
	configuration *models.Configuration
)

type environment struct{}

func (e environment) GetConfiguration() (*models.Configuration, error) {
	if configuration == nil {
		loadConfiguration()
	}
	return configuration, nil
}

func loadConfiguration() {
	commit := getCommit()
	branch := getBranch(commit)
	configuration = &models.Configuration{
		Url:       "localhost",
		Branch:    branch,
		CommitSha: commit,
		Repository: models.Repository{
			Id:     "localhost",
			Name:   "localhost",
			Url:    "localhost",
			Source: getSource(),
		},
		Pipeline: models.Entity{
			Id:   "localhost",
			Name: "localhost",
		},
		Job: models.Entity{
			Id:   "localhost",
			Name: "localhost",
		},
		Run: models.BuildRun{
			BuildId:     "",
			BuildNumber: "",
		},
		Runner: models.Runner{
			Id:   "localhost",
			Name: "localhost",
			OS:   "localhost",
		},
		Environment:   enums.Localhost,
		PipelinePaths: []string{},
		ScmId:         "localhost",
		Pusher: models.Pusher{
			Username: utils.DetectPusher(),
		},
	}
}

func (e environment) Name() string {
	return "localhost"
}

func (e environment) GetStepLink() string {
	return "localhost"
}

func (e environment) GetBuildLink() string {
	return "localhost"
}

func (e environment) GetFileLineLink(filename string, ref string, line int) string {
	return "localhost"
}

func getCommit() string {
	path, _ := os.Getwd()
	commit, _ := git.GetGitCommit(path)
	return commit
}

func getBranch(commit string) string {
	if branch, ok := os.LookupEnv("OVVERIDE_BRANCH"); ok {
		return branch
	}

	path, _ := os.Getwd()
	branch, _ := git.GetGitBranch(path, commit)
	return branch
}

func getSource() enums.Source {
	source, ok := os.LookupEnv("OVVERIDE_BUILDSYSTEM")
	if ok {
		return enums.Source(source)
	}

	return enums.Localhost
}

func (e environment) IsCurrentEnvironment() bool {
	return false
}
