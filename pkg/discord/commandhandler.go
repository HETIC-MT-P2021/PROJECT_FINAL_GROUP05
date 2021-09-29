package discord

func DiscordCommandHandler(userID string, channelID string, content string) (string, bool) {

	switch content {
	case "ping":
		result := Ping(content)
		return result, true
	case "pong":
		result := Ping(content)
		return result, true
	case "carlos":
		result := Ping(content)
		return result, true
	case "contrecarlos":
		result := Ping(content)
		return result, true
	default:
		result := ""
		return result, false
	}
}
