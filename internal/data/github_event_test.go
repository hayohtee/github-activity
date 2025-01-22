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
		{
			expected: "User \"hayohtee\" created a comment to the commit in the repo example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: commitCommentEvent,
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
				Payload: payload{
					Action: "created",
				},
			},
		},
		{
			expected: "User \"hayohtee\" created a branch or tag in the repo example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: createEvent,
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
		{
			expected: "User \"hayohtee\" deleted a branch or tag in the repo example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: deleteEvent,
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
		{
			expected: "User \"hayohtee\" forks the repo example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: forkEvent,
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
		{
			expected: "User \"hayohtee\" created or updated a wiki page",
			githubEvent: GitHubEvent{
				Type: gollumEvent,
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
			},
		},
		{
			expected: "User \"hayohtee\" created a comment to an issue or pull request in the repo example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: issueCommentEvent,
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Payload: payload{
					Action: "created",
				},
			},
		},
		{
			expected: "User \"hayohtee\" opened an issue in the repo example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: issuesEvent,
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Payload: payload{
					Action: "opened",
				},
			},
		},
		{
			expected: "User \"hayohtee\" was added to the example-user/example-repo repository",
			githubEvent: GitHubEvent{
				Type: memberEvent,
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Payload: payload{
					Action: "added",
				},
			},
		},
	}
}
