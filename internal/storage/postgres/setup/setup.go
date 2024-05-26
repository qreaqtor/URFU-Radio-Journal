package setupst

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnect(user, password, host, port, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
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
