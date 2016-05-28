package messages

import "database/sql"

func getDbConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./tz_msgs.sqlite")
	return db, err
}
