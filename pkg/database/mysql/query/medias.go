package query

const (
	QUERY_CREATE_MEDIA string = `
		INSERT INTO media(discord_url, is_archived, user_id, server_id, created_at, updated_at) 
		VALUES(?, ?, ?, ?, ?, ?);
	`
	QUERY_UPDATE_MEDIA string = `
		UPDATE media 
		SET discord_url = ?, is_archived = ?, user_id = ?, updated_at = ?
		WHERE id = ?;
	`
	QUERY_UPDATE_FIELD_IS_ARCHIVED_MEDIA string = `
		UPDATE media 
		SET is_archived = ?, updated_at = ?
		WHERE id = ?;
	`
)