package utils

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetServerResponse(data *models.ServerCommandsAndMedias) gin.H{
	return gin.H{
		"server_name": data.ServerName,
		"created_at": data.CreatedAt,
		"commands": GetServerCommandsResponse(data),
		"medias": GetServerMediasResponse(data),
	}
}

func GetServerCommandsResponse(data *models.ServerCommandsAndMedias) []gin.H {
	var commands []gin.H
	for _, command := range data.Commands {
		displayedCommand := gin.H{
			"title": command.Title,
			"command": command.Command,
			"is_active": command.IsActive,
		}
		commands = append(commands, displayedCommand)
	}

	return commands
}

func GetServerMediasResponse(data *models.ServerCommandsAndMedias) []gin.H {
	var medias []gin.H
	for _, media := range data.Medias {
		displayedMedia := gin.H{
			"discord_url": media.DiscordUrl,
			"is_archived": media.IsArchived,
			"user_id": media.UserID,
		}
		medias = append(medias, displayedMedia)
	}

	return medias
}