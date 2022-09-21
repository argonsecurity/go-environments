package enums

type Source string

const (
	Github          Source = "github"
	Gitlab          Source = "gitlab"
	Azure           Source = "azure"
	Bitbucket       Source = "bitbucket"
	GithubServer    Source = "github_server"
	AzureServer     Source = "azure_server"
	GitlabServer    Source = "gitlab_server"
	BitbucketServer Source = "bitbucket_server"
	Jenkins         Source = "jenkins"
	Localhost       Source = "localhost"
	CircleCi        Source = "circleCI"
	Unknown         Source = "unknown"
)
