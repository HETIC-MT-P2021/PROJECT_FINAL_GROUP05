package query

const (
	QUERY_CREATE_MEDIA string = `
		INSERT INTO media(discord_url, is_archived, user_id, server_id) VALUES(?, ?, ?, ?);
	`
)