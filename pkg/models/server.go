package models

type Server struct {
	ID 						int 			`json:"id"`
	ServerID 			string 		`json:"server_id"`
	ServerName 		string 		`json:"server_name"`
  CreatedAt 		string 		`json:"created_at"`
  UpdatedAt 		string 		`json:"updated_at"`
}

type ServerCommandsAndMedias struct {
	ID 						int 			`json:"id"`
	ServerID 			string 		`json:"server_id"`
	ServerName 		string 		`json:"server_name"`
  CreatedAt 		string 		`json:"created_at"`
  UpdatedAt 		string 		`json:"updated_at"`
	Commands 			[]Command `json:"commands"`
	Medias 				[]Media 	`json:"medias"`
}