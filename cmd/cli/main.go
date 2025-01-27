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

	pageFlag := flag.NewFlagSet("", flag.ExitOnError)

	page := pageFlag.Int("page", 1, "The starting page for github events")
	pageSize := pageFlag.Int("page-size", 50, "The  maximum number of github events to fetch per page")
	pageFlag.Parse(os.Args[2:])

	username := os.Args[1]

	events, err := fetchGithubEvents(username, *page, *pageSize)
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
func fetchGithubEvents(username string, page, pageSize int) ([]data.GitHubEvent, error) {
	var githubEvents []data.GitHubEvent

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	apiUrl := fmt.Sprintf("https://api.github.com/users/%s/events?page=%d&per_page=%d", username, page, pageSize)

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&githubEvents); err != nil {
		return nil, err
	}

	return githubEvents, nil
}

// printGithubEvents prints a list of GitHub events to the standard output.
// Each event is printed on a new line prefixed with a hyphen.
//
// Parameters:
//
//	events ([]data.GitHubEvent): A slice of GitHubEvent objects to be printed.
func printGithubEvents(events []data.GitHubEvent) {
	fmt.Println("Output:")
	for _, event := range events {
		fmt.Printf("- %s\n", event)
	}
}
