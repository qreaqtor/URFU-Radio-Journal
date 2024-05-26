package setupst

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnect(user, password, host, dbName string, port int) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to PostgreSQL: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while trying to ping PostgreSQL: %v", err)
	}

	return db, nil
}
