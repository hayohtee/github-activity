package data

// repo is a struct containing the id and the name of the github repository.
type repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
