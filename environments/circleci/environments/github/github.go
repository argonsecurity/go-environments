package github

import (
	"context"
	"github.com/google/go-github/v44/github"
	"golang.org/x/oauth2"
	"os"
)

func GetPullRequestTargetBranch(owner, repo string, prNumber int) (string, error) {
	client, err := newGithubClient(os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		return "", err
	}
	pr, _, err := client.PullRequests.Get(context.Background(), owner, repo, prNumber)
	if err != nil {
		return "", err
	}

	return *pr.Base.Ref, nil
}

func newGithubClient(token string) (*github.Client, error) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc), nil
}
