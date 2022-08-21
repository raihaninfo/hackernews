package models

import "time"

type Post struct {
	ID           int       `db:"id,omitempty"`
	Title        string    `db:"title"`
	Url          string    `db:"url"`
	CreatedAt    time.Time `db:"created_at"`
	UserID       int       `db:"user_id"`
	Votes        int       `db:"votes,omitempty"`
	UserName     string    `db:"user_name,omitempty"`
	CommentCount int       `db:"comment_count,omitempty"`
	TotalRecords int       `db:"total_records,omitempty"`
}
