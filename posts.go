package ghost

import (
	"fmt"
	"time"
)

const (
	// Post statuses
	PostStatusPublished = "published"
	PostStatusDraft     = "draft"
)

// Post represents a post in Ghost
type Post struct {
	ID          int        `json:"id"`
	UUID        string     `json:"uuid"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Markdown    string     `json:"markdown"`
	HTML        string     `json:"html"`
	Featured    intBool    `json:"featured"`
	Page        intBool    `json:"page"`
	Status      string     `json:"status"`
	Language    string     `json:"language"`
	PublishedAt *time.Time `json:"published_at"`
	PublishedBy int        `json:"published_by"`
	CreatedAt   *time.Time `json:"created_at"`
	CreatedBy   int        `json:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   int        `json:"updated_by"`
}

// intBool parses an integer as a boolean, where 0 is false and 1 is true
type intBool bool

func (bit *intBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "1" || s == "true" {
		*bit = true
	} else if s == "0" || s == "false" {
		*bit = false
	} else {
		return fmt.Errorf("intBool unmarshal error: invalid input %s", s)
	}
	return nil
}
