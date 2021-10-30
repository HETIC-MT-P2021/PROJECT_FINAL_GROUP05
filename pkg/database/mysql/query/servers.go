package query

const (
	QUERY_CREATE_SERVER string = `
		INSERT INTO srv(id, server_name, created_at) VALUES(?, ?, ?);
	`
	QUERY_FIND_SERVERS string = `
		SELECT id, server_name, created_at
		FROM srv;
	`
	QUERY_FIND_SERVER_WITH_MADIA_AND_CMDS string = `
		SELECT server_name, created_at, 
					 command.title, command.command, command.is_active,
					 media.discord_url, media.is_archived, media.user_id
		FROM srv
		JOIN command ON srv.id = command.server_id
		JOIN media ON srv.id = media.server_id
		WHERE srv.id = ?;
	`
)