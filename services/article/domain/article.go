package domain

import "time"

type Article struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Url           string    `json:"url"`
	BookmarkCount uint      `json:"bookmark_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
