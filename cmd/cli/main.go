package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hayohtee/github-activity/internal/data"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no argument provided")
		os.Exit(1)
	}

	page := flag.Int("page", 30, "The starting page for github events")
	pageSize := flag.Int("page-size", 50, "The  maximum number of github events to fetch per page")

	username := os.Args[1]

	events, err := fetchGithubEvents(username)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	printGithubEvents(events)
}

// fetchGithubEvents retrieves a list of GitHub events for a specified user.
// It sends a GET request to the GitHub API to fetch the events.
//
// Parameters:
//   - username: The GitHub username for which to fetch events.
//   - page: The page number of the results to retrieve.
//   - pageSize: The number of results per page.
//
// Returns:
//   - A slice of GitHubEvent objects containing the fetched events.
//   - An error if the request fails or the response cannot be decoded.
