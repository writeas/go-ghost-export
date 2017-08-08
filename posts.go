package ghost

import "time"

const (
	// Post statuses
	// TODO: add other statuses
	PostStatusPublished = "published"
)

// Post represents a post in Ghost
type Post struct {
	ID          int       `json:"id"`
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Markdown    string    `json:"markdown"`
	HTML        string    `json:"html"`
	Featured    bool      `json:"featured"`
	Page        bool      `json:"page"`
	Status      string    `json:"status"`
	Language    string    `json:"language"`
	PublishedAt time.Time `json:"published_at"`
	PublishedBy int       `json:"published_by"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   int       `json:"updated_by"`
}
