package github

type GithubOwner struct {
	Login string `json:"login"`
}

type GithubRepository struct {
	Id    int         `json:"id"`
	Name  string      `json:"name"`
	Owner GithubOwner `json:"owner"`
}

type GithubAuthor struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type GithubCommit struct {
	Id        string       `json:"id"`
	Message   string       `json:"message"`
	Timestamp string       `json:"timestamp"`
	Author    GithubAuthor `json:"author"`
	Url       string       `json:"url"`
}

type GithubSender struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}

type GithubPayload struct {
	Repository GithubRepository `json:"repository"`
	Sender     GithubSender     `json:"sender"`
	Commits    []GithubCommit   `json:"commits"`
}
