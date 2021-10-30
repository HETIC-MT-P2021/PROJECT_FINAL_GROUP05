package models

type Server struct {
	ID 						string 		`json:"id"`
	ServerName 		string 		`json:"server_name"`
  CreatedAt 		string 		`json:"created_at"`
}

type ServerCommandsAndMedias struct {
	ServerName 		string 		`json:"server_name"`
  CreatedAt 		string 		`json:"created_at"`
	Commands 			[]Command `json:"commands"`
	Medias 				[]Media 	`json:"medias"`
}