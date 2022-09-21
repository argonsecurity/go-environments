package azure

import (
	_ "embed"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/argonsecurity/go-utils/environments/enums"
	"github.com/argonsecurity/go-utils/environments/environments/utils"
	"github.com/argonsecurity/go-utils/environments/environments/utils/git"
	"github.com/argonsecurity/go-utils/environments/models"
	"github.com/argonsecurity/go-utils/logger"
	schemavalidator "github.com/argonsecurity/go-utils/schema-validator"
)

const (
	taskInstanceIDEnv = "SYSTEM_TASKINSTANCEID"
	projectIDEnv      = "SYSTEM_TEAMPROJECTID"
	definitionIDEnv   = "SYSTEM_DEFINITIONID"
	projectNameEnv    = "SYSTEM_TEAMPROJECT"
	jobIDEnv          = "SYSTEM_JOBID"
	jobNameEnv        = "SYSTEM_JOBDISPLAYNAME"
	repositoryPathEnv = "BUILD_SOURCESDIRECTORY"
	branchEnv         = "BUILD_SOURCEBRANCH"
	userEmailEnv      = "BUILD_REQUESTEDFOREMAIL"
	pipelineNameEnv   = "BUILD_DEFINITIONNAME"
	buildIDEnv        = "BUILD_BUILDID"
	buildNumberEnv    = "BUILD_BUILDNUMBER"
	endpointURLEnv    = "SYSTEM_TASKDEFINITIONSURI"
	collectionUriEnv  = "SYSTEM_COLLECTIONURI"
	commitShaEnv      = "BUILD_SOURCEVERSION"

	collectionIdEnv = "SYSTEM_COLLECTIONID"

	pullRequestIdEnv           = "SYSTEM_PULLREQUEST_PULLREQUESTID"
	pullRequestSourceBranchEnv = "SYSTEM_PULLREQUEST_SOURCEBRANCH"
	pullRequestTargetBranchEnv = "SYSTEM_PULLREQUEST_TARGETBRANCH"

	repositoryIdEnv   = "BUILD_REPOSITORY_ID"
	repositoryNameEnv = "BUILD_REPOSITORY_NAME"
	repositoryUriEnv  = "BUILD_REPOSITORY_URI"

	usernameEnv = "BUILD_QUEUEDBY"

	agentIDEnv             = "AGENT_ID"
	agentNameEnv           = "AGENT_NAME"
	agentOSArchitectureEnv = "AGENT_OSARCHITECTURE"
	agentOSEnv             = "AGENT_OS"
	imageOSEnv             = "ImageOS"

	buildReasonEnv    = "BUILD_REASON"
	DetectionVariable = "BUILD_BUILDID"

	azureDevopsApiUrlEnv  = "ENDPOINT_URL_SYSTEMVSSCONNECTION"
	azurePullRequestEvent = "PullRequest"
)

var (
	//go:embed azure-pipelines.schema.json
	azurePipelinesSchema []byte

	// Azure environment
	Azure         = environment{}
	configuration *models.Configuration
	baseUrlRegex  = regexp.MustCompile(`https:\/\/[\w.]*(dev.azure.com|vsassets.io|vsassets.io|msauth.net|msftauth.net|visualstudio.com|azure.net|microsoft.com|azurecomcdn.azureedge.net|live.com|microsoftonline.com|management.azure.com|sharepointonline.com|.windows.net|azureedge.net)`)
)

type environment struct{}

func (e environment) GetConfiguration() (*models.Configuration, error) {
	if configuration == nil {
		if err := loadConfiguration(); err != nil {
			return nil, err
		}
	}
	return configuration, nil
}

func loadConfiguration() error {
	repoPath := os.Getenv(repositoryPathEnv)
	repoUrl := os.Getenv(repositoryUriEnv)
	cloneUrl, err := git.GetGitRemoteURL(repoPath)

	if err != nil || cloneUrl == "" {
		cloneUrl = fmt.Sprintf("%s.git", repoUrl)
	}

	source := getSource()

	userName := os.Getenv(usernameEnv)
	if userName == "" {
		userName = utils.DetectPusher()
	}

	configuration = &models.Configuration{
		Url:       os.Getenv(endpointURLEnv),
		SCMApiUrl: os.Getenv(azureDevopsApiUrlEnv),
		LocalPath: repoPath,
		Branch:    getBranch(),
		ProjectId: os.Getenv(projectIDEnv),
		CommitSha: os.Getenv(commitShaEnv),
		Organization: models.Entity{
			Id:   os.Getenv(collectionIdEnv),
			Name: getOrganizationName(os.Getenv(collectionUriEnv)),
		},
		Repository: models.Repository{
			Id:       os.Getenv(repositoryIdEnv),
			Name:     os.Getenv(repositoryNameEnv),
			Url:      repoUrl,
			CloneUrl: cloneUrl,
			Source:   source,
		},
		Pusher: models.Pusher{
			Username: userName,
			Email:    os.Getenv(userEmailEnv),
		},
		Job: models.Entity{
			Id:   os.Getenv(jobNameEnv),
			Name: os.Getenv(jobNameEnv),
		},
		Pipeline: models.Entity{
			Id:   fmt.Sprintf("%s-%s", os.Getenv(projectIDEnv), os.Getenv(definitionIDEnv)),
			Name: os.Getenv(pipelineNameEnv),
		},
		Run: models.BuildRun{
			BuildId:     os.Getenv(buildIDEnv),
			BuildNumber: os.Getenv(buildNumberEnv),
		},
		Runner: models.Runner{
			Id:           os.Getenv(agentIDEnv),
			Name:         os.Getenv(agentNameEnv),
			OS:           os.Getenv(agentOSEnv),
			Distribution: os.Getenv(imageOSEnv),
			Architecture: os.Getenv(agentOSArchitectureEnv),
		},
		PullRequest: models.PullRequest{
			Id: os.Getenv(pullRequestIdEnv),
			SourceRef: models.Ref{
				Branch: os.Getenv(pullRequestSourceBranchEnv),
			},
			TargetRef: models.Ref{
				Branch: os.Getenv(pullRequestTargetBranchEnv),
			},
		},
		PipelinePaths: getPipelinePaths(repoPath),
		Environment:   source,
	}
	return nil
}

func (e environment) GetStepLink() string {
	return fmt.Sprintf("%s%s/_build/results?buildId=%s&view=logs&j=%s&t=%s", os.Getenv(endpointURLEnv), os.Getenv(projectNameEnv),
		os.Getenv(buildIDEnv), os.Getenv(jobIDEnv), os.Getenv(taskInstanceIDEnv))
}

func (e environment) GetBuildLink() string {
	return fmt.Sprintf("%s%s/_build?definitionId=%s&_a=summary", os.Getenv(endpointURLEnv), os.Getenv(projectNameEnv), os.Getenv(definitionIDEnv))
}

func (e environment) GetFileLineLink(filePath string, branchName string, lineNumber int) string {
	if lineNumber != 0 {
		return fmt.Sprintf("%s_git/%s?path=%s&version=GB%s&line=%d&lineEnd=%d&lineStartColumn=1&lineEndColumn=1&lineStyle=plain&_a=contents",
			os.Getenv(endpointURLEnv),
			os.Getenv(repositoryNameEnv),
			strings.Replace(filePath, "/", "%2F", -1),
			branchName,
			lineNumber,
			lineNumber+1,
		)
	}
	return fmt.Sprintf("%s_git/%s?path=%s&version=GB%s&_a=contents",
		os.Getenv(endpointURLEnv),
		os.Getenv(repositoryNameEnv),
		strings.Replace(filePath, "/", "%2F", -1),
		branchName,
	)
}

func getBranch() string {
	if os.Getenv(buildReasonEnv) == azurePullRequestEvent {
		return os.Getenv(pullRequestSourceBranchEnv)
	}

	return os.Getenv(branchEnv)
}

func getSource() enums.Source {
	if baseUrlRegex.MatchString(os.Getenv(collectionUriEnv)) {
		return enums.Azure
	}
	return enums.AzureServer
}

func getOrganizationName(collectionURI string) string {
	parsedURL, err := url.ParseRequestURI(collectionURI)
	if err != nil {
		return ""
	}

	split := strings.Split(strings.TrimSuffix(parsedURL.Path, "/"), "/")
	return split[len(split)-1]
}

func (e environment) IsCurrentEnvironment() bool {
	_, isExist := os.LookupEnv("BUILD_BUILDID")
	return isExist
}

func (e environment) Name() string {
	return "azure"
}

func getPipelinePaths(rootDir string) []string {
	paths := make([]string, 0)

	filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if (filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml") && isAzurePipeline(path, azurePipelinesSchema) {
			paths = append(paths, path)
		}
		return nil
	})

	return paths
}

func isAzurePipeline(filePath string, schema []byte) bool {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		logger.Warnf("Failed to read yml file %s", filePath)
		return false
	}

	err = schemavalidator.ValidateYaml(fileData, schema)
	return err == nil
}
