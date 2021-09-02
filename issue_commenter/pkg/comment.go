package pkg

type Config struct {
	Platform  string
	EventType string
	Payload   string
	Token     string
	Message   string
}

// Commenter takes a PR request and comments on it, if the previous
// comment had a trigger key word.
type Commenter struct {
	Config
}

// NewCommenter creates a new Commenter.
func NewCommenter(cfg Config) Commenter {
	return Commenter{Config: cfg}
}

// Comment sends a message to a PR.
func (c Commenter) Comment() error {
	return nil
}
