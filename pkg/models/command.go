package models

type Command struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Command string `json:"command"`
	IsChecked bool `json:"is_checked"`
}