package query

const (
	QUERY_CREATE_SERVER string = `
		INSERT INTO srv(id, server_name, created_at) VALUES(?, ?, ?);
	`
	QUERY_FIND_SERVERS string = `
		SELECT id, server_name, created_at
		FROM srv;
	`
	QUERY_FIND_SERVER string = `
		SELECT id, server_name, created_at
		FROM srv
		WHERE id = ?;
	`
	QUERY_FIND_SERVER_MEDIAS string = `
		SELECT media.discord_url, media.is_archived, media.user_id
		FROM srv
		JOIN media ON srv.id = media.server_id
		WHERE srv.id = ?;
	`
	QUERY_FIND_SERVER_COMMANDS string = `
		SELECT command.title, command.command, command.is_active
		FROM srv
		JOIN command ON srv.id = command.server_id
		WHERE srv.id = ?;
	`
)