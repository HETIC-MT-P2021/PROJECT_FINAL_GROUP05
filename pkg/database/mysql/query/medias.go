package query

const (
	QUERY_CREATE_MEDIA string = `
		INSERT INTO media(discord_url, is_archived, user_id, server_id) VALUES(?, ?, ?, ?);
	`
	QUERY_UPDATE_MEDIA string = `
		UPDATE media 
		SET discord_url = ?, is_archived = ?, user_id = ?
		WHERE id = ?;
	`
	QUERY_UPDATE_FIELD_IS_ARCHIVED_MEDIA string = `
		UPDATE media 
		SET is_archived = ?
		WHERE id = ?;
	`
)