package data

import "testing"

type testItem struct {
	expected    string
	githubEvent GitHubEvent
}

func TestGitHubEvent(t *testing.T) {
	testTable := []testItem{
		{
			expected: "User \"hayohtee\" starred the example-user/example-repo",
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
			expected: "User \"hayohtee\" created a comment to the commit in the example-user/example-repo",
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
			expected: "User \"hayohtee\" created a branch or tag in the example-user/example-repo",
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
			expected: "User \"hayohtee\" deleted a branch or tag in the example-user/example-repo",
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
			expected: "User \"hayohtee\" forks the example-user/example-repo",
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
			expected: "User \"hayohtee\" created a comment to an issue or pull request in the example-user/example-repo",
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
			expected: "User \"hayohtee\" opened an issue in the example-user/example-repo",
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
			expected: "User \"hayohtee\" was added to the example-user/example-repo",
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
		{
			expected: "example-user/example-repo was made public",
			githubEvent: GitHubEvent{
				Type: publicEvent,
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
			},
		},
		{
			expected: "User \"hayohtee\" opened a pull request in example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: pullRequestEvent,
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
				Payload: payload{
					Action: "opened",
				},
			},
		},
		{
			expected: "User \"hayohtee\" created a review in the pull request to example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: pullRequestReviewEvent,
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
			expected: "User \"hayohtee\" created a comment in a review to the pull request in example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: pullRequestReviewCommentEvent,
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
			expected: "User \"hayohtee\" resolved the comment thread in the review to the pull request in example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: pullRequestReviewThreadEvent,
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
				Payload: payload{
					Action: "resolved",
				},
			},
		},
		{
			expected: "User \"hayohtee\" pushed 3 commits to the example-user/example-repo branch or tag",
			githubEvent: GitHubEvent{
				Type: pushEvent,
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
				Payload: payload{
					Size: 3,
				},
			},
		},
		{
			expected: "User \"hayohtee\" published a release in example-user/example-repo",
			githubEvent: GitHubEvent{
				Type: releaseEvent,
				Repo: repo{
					ID:   0,
					Name: "example-user/example-repo",
				},
				Actor: actor{
					Login:        "hayohtee",
					DisplayLogin: "hayohtee",
				},
				Payload: payload{
					Action: "published",
				},
			},
		},
		{
			expected: "User \"hayohtee\" created a sponsorship listing",
			githubEvent: GitHubEvent{
				Type: sponsorshipEvent,
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
	}
}
