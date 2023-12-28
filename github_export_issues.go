package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

var pretty = flag.Bool("pretty", false, "Pretty Print JSON")
var authToken = flag.String("auth_token", "", "Auth token for better rate limits")
var state = flag.String("state", "all", "Issue state")
var issueID = flag.Int("issue-id", -1, "Only export 1 issue")

func main() {
	flag.Parse()
	ctx := context.Background()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("Usage: %s <project_url>", os.Args[0])
	}

	projectURL := args[0]

	u, err := url.Parse(projectURL)
	if err != nil {
		log.Fatalf("URL parse error: %s\n", err)
	}

	if !strings.HasSuffix(u.Hostname(), "github.com") {
		log.Fatalf("Expected a github.com domain")
	}

	parts := strings.Split(u.Path, "/")

	if len(parts) < 3 {
		log.Fatalf("Failed to find owner/repo in URL path")
	}

	owner := parts[1]
	repo := parts[2]

	var hc *http.Client
	if *authToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: *authToken},
		)
		hc = oauth2.NewClient(ctx, ts)
	}

	client := github.NewClient(hc)

	enc := json.NewEncoder(os.Stdout)
	if *pretty {
		enc.SetIndent("", "  ")
	}

	if *issueID > 0 {
		issue, _, err := client.Issues.Get(ctx, owner, repo, *issueID)
		if err != nil {
			log.Fatalf("Fetch issue %d for %s/%s failed: %s", *issueID, owner, repo, err)
		}
		allComments, err := fetchComments(ctx, client, owner, repo, issue)
		if err != nil {
			log.Fatalf("fetch comments for issue %s/%s failed: %s", owner, repo, err)
		}

		enc.Encode(Issue{
			Issue:    *issue,
			Comments: allComments,
		})

		os.Exit(0)

	}

	var issueOps github.IssueListByRepoOptions
	issueOps.State = *state
	for {
		issues, issuesResp, err := client.Issues.ListByRepo(ctx, owner, repo, &issueOps)
		if err != nil {
			log.Fatalf("Fetch issues for %s/%s failed: %s", owner, repo, err)
		}
		for _, issue := range issues {

			allComments, err := fetchComments(ctx, client, owner, repo, issue)
			if err != nil {
				log.Fatalf("fetch comments for issue %s/%s failed: %s", owner, repo, err)
			}

			enc.Encode(Issue{
				Issue:    *issue,
				Comments: allComments,
			})
		}

		if issuesResp.NextPage == 0 {
			break
		}

		issueOps.Page = issuesResp.NextPage
	}
}

func fetchComments(ctx context.Context, client *github.Client, owner, repo string, issue *github.Issue) ([]github.IssueComment, error) {
	var commentOps github.IssueListCommentsOptions
	var allComments []github.IssueComment

	for {
		comments, commentsResp, err := client.Issues.ListComments(ctx, owner, repo, issue.GetNumber(), &commentOps)
		if err != nil {
			return nil, err
		}

		for _, comment := range comments {
			allComments = append(allComments, *comment)
		}

		if commentsResp.NextPage == 0 {
			break
		}
		commentOps.Page = commentsResp.NextPage
	}

	return allComments, nil

}

type Issue struct {
	Issue    github.Issue          `json:"issue"`
	Comments []github.IssueComment `json:"comments"`
}
