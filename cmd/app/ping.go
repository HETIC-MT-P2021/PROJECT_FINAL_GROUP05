package main

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func ping(Content string) string {

	if Content == "ping" || Content == "carlos" {
		return "https://cdn.discordapp.com/attachments/694885042400264266/892463464591790190/carlos.jpg"
	}

	// If the message is "pong" reply with "Ping!"
	if Content == "pong" || Content == "contrecarlos" {
		return "https://media.discordapp.net/attachments/694885042400264266/892464389855256586/solrac.png"
	}
	return ""
}
