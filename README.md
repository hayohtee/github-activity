# github-activity
A simple command line interface (CLI) to fetch the recent activity of a GitHub 
user and display it in the terminal.\
Solution to [GitHub User Activity](https://roadmap.sh/projects/github-user-activity/)
on [Roadmap.sh](https://roadmap.sh/).

## Supported Event Types
- CommitCommentEvent
- CreateEvent
- DeleteEvent
- ForkEvent
- GollumEvent
- IssueCommentEvent
- IssuesEvent
- MemberEvent
- PublicEvent
- PullRequestEvent
- PullRequestReviewEvent
- PullRequestReviewCommentEvent
- PullRequestReviewThreadEvent
- PushEvent
- ReleaseEvent
- SponsorshipEvent
- WatchEvent

## Building from source
Ensure the GO SDK is installed
1. Clone the repository
   ```bash
   git https://github.com/hayohtee/github-activity.git
   ```
3. Change into the project directory
   ```bash
   cd github-activity
   ```
4. Compile
   ```bash
   go build -o github-activity ./cmd/cli
   ```

## Usage
```bash
# Fetch the recent github activities, default page is 1 and the maximum page size is 30
./github-activity <username>

# Use --page flag to specify starting page in INT and --page-size to specify the maximum
# number of github activity to return also in INT.
./github-activity <username> --page=<value> --page-size=<value>
```
