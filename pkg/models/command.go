package models

type Command struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Command string `json:"command"`
	IsActive bool `json:"is_active"`
	ServerID string `json:"server_id"`
}