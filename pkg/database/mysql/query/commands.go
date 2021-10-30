package query

const (
	QUERY_CREATE_COMMAND string = `
		INSERT INTO command(title, command, is_active, server_id) VALUES(?, ?, ?, ?);
	`
	QUERY_FIND_COMMANDS string = `
		SELECT title, command, is_active, server_id
		FROM command;
	`
	QUERY_UPDATE_COMMAND string = `
		UPDATE command 
		SET title = ?, command = ?, is_active = ?, server_id = ?
		WHERE id = ?;
	`
	QUERY_UPDATE_FIELD_IS_ACTIVE_COMMAND string = `
		UPDATE command 
		SET is_active = ?
		WHERE id = ?;
	`
)