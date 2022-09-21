package localhost

import (
	"github.com/argonsecurity/go-environments/enums"
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
	configuration = &models.Configuration{
		Url: "localhost",
		Repository: models.Repository{
			Id:     "localhost",
			Name:   "localhost",
			Url:    "localhost",
			Source: enums.Localhost,
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
			BuildId:     "localhost",
			BuildNumber: "localhost",
		},
		Runner: models.Runner{
			Id:   "localhost",
			Name: "localhost",
			OS:   "localhost",
		},
		Environment:   enums.Localhost,
		PipelinePaths: []string{},
		ScmId:         "localhost",
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

func (e environment) IsCurrentEnvironment() bool {
	return false
}
