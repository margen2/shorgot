package models

import "time"

type Link struct {
	ID        uint64    `json:"Id,omitempty"`
	AuthorID  uint64    `json:"authorId,omitempty"`
	Clicks    uint64    `json:"clicks"`
	Target    string    `json:"target,omitempty"`
	Shortened string    `json:"shortened,omitempty"`
	CreatedOn time.Time `json:"createdon,omitempty"`
}
