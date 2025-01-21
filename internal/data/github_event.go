package data

import "time"

const (
	watchEvent = "WatchEvent"
	commitCommentEvent = "CommitCommentEvent"
	createEvent = "CreateEvent"
)

// repo is a struct containing the id and the name of the github repository.
type repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// payload is a struct containing the size and distinct size of the commits.
type payload struct {
	Size         int `json:"size"`
	DistinctSize int `json:"distinct_size"`
}

// GitHubEvent is a struct containing all the information about a particular github event.
type GitHubEvent struct {
	Type      string    `json:"type"`
	Repo      repo      `json:"repo"`
	Payload   payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}
