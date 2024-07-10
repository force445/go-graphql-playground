package model

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}
