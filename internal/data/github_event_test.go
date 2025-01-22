package data

import "testing"

type testItem struct {
	expected    string
	githubEvent GitHubEvent
}

func TestGitHubEvent(t *testing.T) {
	testTable := []testItem{
		{
			expected: "User \"hayohtee\" stars the repo example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: watchEvent,
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
			},
		},
	}
}
