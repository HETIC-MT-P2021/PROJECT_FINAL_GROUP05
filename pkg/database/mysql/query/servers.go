package query

const (
	QUERY_CREATE_SERVER string = `
		INSERT INTO srv(id, server_name, created_at) VALUES(?, ?, ?);
	`
	QUERY_FIND_SERVERS string = `
		SELECT id, server_name, created_at 
		FROM srv;
	`
)