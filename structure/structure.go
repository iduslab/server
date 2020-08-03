package structure

// Structure of config.json
type Config struct {
	Token   string `json:"token"`
	OwnerID string `json:"ownerID"`
	Prefix  string `json:"prefix"`
}

// Structure of IdeaNote
type Note struct {
	BoxNum     int
	Author     string
	AuthorName string
	text       string
}

// Structure of Idea Box
type Box struct {
	ID   int
	Text string
}
