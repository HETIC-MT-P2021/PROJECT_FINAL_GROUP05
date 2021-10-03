package query

const (
	QUERY_CREATE_COMMAND string = "INSERT INTO commands(title, command, is_checked) VALUES(?, ?, ?)"
	QUERY_FIND_COMMANDS string = "SELECT title, command, is_checked FROM commands"
)