package models

type Media struct {
	ID 					int 	 `json:"id"`
  DiscordUrl 	string `json:"discord_url"`
  IsArchived 	bool	 `json:"is_archived"`
  UserID 			string `json:"user_id"`
  ServerID 		string `json:"server_id"`
  CreatedAt 		string 		`json:"created_at"`
  UpdatedAt 		string 		`json:"updated_at"`
}