package data

import (
	"fmt"
	"time"
)

const (
	watchEvent                    = "WatchEvent"
	commitCommentEvent            = "CommitCommentEvent"
	createEvent                   = "CreateEvent"
	deleteEvent                   = "DeleteEvent"
	forkEvent                     = "ForkEvent"
	gollumEvent                   = "GollumEvent"
	issueCommentEvent             = "IssueCommentEvent"
	issuesEvent                   = "IssuesEvent"
	memberEvent                   = "MemberEvent"
	publicEvent                   = "PublicEvent"
	pullRequestEvent              = "PullRequestEvent"
	pullRequestReviewEvent        = "PullRequestReviewEvent"
	pullRequestReviewCommentEvent = "PullRequestReviewCommentEvent"
	pullRequestReviewThreadEvent  = "PullRequestReviewThreadEvent"
	pushEvent                     = "PushEvent"
	releaseEvent                  = "ReleaseEvent"
	sponsorshipEvent              = "SponsorshipEvent"
)

// repo is a struct containing the id and the name of the github repository.
type repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// payload is a struct containing the size and distinct size of the commits.
type payload struct {
	Size         int    `json:"size"`
	DistinctSize int    `json:"distinct_size"`
	Action       string `json:"action"`
}

// actor is a struct containing the login and display_login name of an actor.
type actor struct {
	login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
}

// GitHubEvent is a struct containing all the information about a particular github event.
type GitHubEvent struct {
	Type      string    `json:"type"`
	Repo      repo      `json:"repo"`
	Payload   payload   `json:"payload"`
	Public    bool      `json:"public"`
	Actor     actor     `json:"actor"`
	CreatedAt time.Time `json:"created_at"`
}

func (g GitHubEvent) String() string {
	switch g.Type {
	case watchEvent:
		return fmt.Sprintf("")
	default:
		return "Unsupported activity"
	}
}