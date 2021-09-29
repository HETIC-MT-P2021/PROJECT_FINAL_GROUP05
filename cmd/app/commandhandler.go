package main

func discordcommandHandler(userID string, channelID string, content string) (string, bool) {

	switch content {
	case "ping":
		result := ping(content)
		return result, true
	case "pong":
		result := ping(content)
		return result, true
	case "carlos":
		result := ping(content)
		return result, true
	case "contrecarlos":
		result := ping(content)
		return result, true
	default:
		result := ""
		return result, false
	}
}
