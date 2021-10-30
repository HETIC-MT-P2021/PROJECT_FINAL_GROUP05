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
		SELECT server_name, created_at, command.title, command.command, command.is_active
		FROM srv
		JOIN command ON srv.id = command.server_id
		WHERE srv.id = ?;
	`
)