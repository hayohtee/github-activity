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
	Login        string `json:"login"`
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
		return fmt.Sprintf("User %q %s the repository %q", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case commitCommentEvent:
		return fmt.Sprintf("User %q %s commit comment in the repository %q", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case createEvent:
		return fmt.Sprintf("Created a branch in %s", g.Repo.Name)
	case deleteEvent:
		return fmt.Sprintf("Deleted a branch in %s", g.Repo.Name)
	case forkEvent:
		return fmt.Sprintf("User %q fork the %s", g.Actor.DisplayLogin, g.Repo.Name)
	case gollumEvent:
		return fmt.Sprintf("User %q created or updated a wiki page", g.Actor.DisplayLogin)
	case issueCommentEvent:
		return fmt.Sprintf("User %q %s an issue comment in %s", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case issuesEvent:
		return fmt.Sprintf("User %q %s an issue in %s", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case memberEvent:
		return fmt.Sprintf("User %q %s an invitation to %s", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case publicEvent:
		return fmt.Sprintf("User %q made the %s public", g.Actor.DisplayLogin, g.Repo.Name)
	case pullRequestEvent:
		return fmt.Sprintf("User %q %s a pull request in %s", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case pullRequestReviewEvent:
		return fmt.Sprintf("User %q %s a pull request review in %s", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case pullRequestReviewCommentEvent:
		return fmt.Sprintf("User %q %s a pull request review comment in %s", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case pullRequestReviewThreadEvent:
		return fmt.Sprintf("User %q %s a comment thread on a pull request in %s", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case pushEvent:
		return fmt.Sprintf("User %q pushed %d commits to %s", g.Actor.DisplayLogin, g.Payload.Size, g.Repo.Name)
	case releaseEvent:
		return fmt.Sprintf("User %q %s a release in %s", g.Actor.DisplayLogin, g.Payload.Action, g.Repo.Name)
	case sponsorshipEvent:
		return fmt.Sprintf("User %q %s a sponsorship listing", g.Actor.DisplayLogin, g.Payload.Action)
	default:
		return "Unsupported activity"
	}
}
