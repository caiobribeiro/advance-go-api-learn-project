package models

type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Cotent  string `json:"content"`
	UserId  string `json:"user_id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
