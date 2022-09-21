package bitbucket

type BitbucketUser struct {
	Uuid        string `json:"uuid"`
	DisplayName string `json:"display_name"`
}

type BitbucketProject struct {
	Uuid string `json:"uuid"`
}

type BitbucketRepository struct {
	Name    string               `json:"name"`
	Uuid    string               `json:"uuid"`
	Links   BitbucketLinksEntity `json:"links"`
	Project BitbucketProject     `json:"project"`
}

type BitbucketCommit struct {
	Hash string `json:"hash"`
}

type BitbucketLinkEntity struct {
	Href string `json:"href"`
}

type BitbucketLinksEntity struct {
	Html BitbucketLinkEntity `json:"html"`
}

type BitbucketPullRequestEntity struct {
	Commit BitbucketCommit `json:"commit"`
}

type BitbucketPullRequest struct {
	Id          string                     `json:"id"`
	Title       string                     `json:"title"`
	Source      BitbucketPullRequestEntity `json:"source"`
	Destination BitbucketPullRequestEntity `json:"destination"`
	Links       BitbucketLinksEntity       `json:"links"`
	Author      BitbucketUser              `json:"author"`
}

type BitbucketPayload struct {
	Repository  BitbucketRepository  `json:"repository"`
	Actor       BitbucketUser        `json:"actor"`
	PullRequest BitbucketPullRequest `json:"pullrequest"`
}
