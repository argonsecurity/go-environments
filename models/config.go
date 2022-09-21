package models

import "github.com/argonsecurity/go-environments/enums"

type Runner struct {
	Id           string
	Name         string
	OS           string
	Distribution string
	Architecture string
}

type Entity struct {
	Id   string
	Name string
}

type Ref struct {
	Sha    string
	Branch string
}

type PullRequest struct {
	Id        string
	SourceRef Ref
	TargetRef Ref
	Url       string
}

type Repository struct {
	Id       string
	Name     string
	Url      string
	CloneUrl string
	Source   enums.Source
}

type Configuration struct {
	Url             string
	SCMApiUrl       string
	Builder         string
	LocalPath       string
	CommitSha       string
	BeforeCommitSha string
	Branch          string
	ProjectId       string
	Job             Entity
	Run             BuildRun
	Pipeline        Entity
	Runner          Runner
	Repository      Repository
	PullRequest     PullRequest
	Commits         []Commit
	Organization    Entity
	Pusher          Pusher
	PipelinePaths   []string
	Environment     enums.Source
	ScmId           string
}

type Author struct {
	Email    string
	Name     string
	Username string
}

type Pusher struct {
	Entity
	Username string
	Email    string
}

type BuildRun struct {
	BuildId     string
	BuildNumber string
}

type Commit struct {
	Id         string
	Message    string
	CommitDate string
	Url        string
	Author     Author
}
