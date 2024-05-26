package setupst

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx"
)

func GetConnect(user, password, host, port, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host+":"+port,
		port,
		dbName,
	)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close()
	if err != nil {
		return nil, fmt.Errorf("error while connecting to PostgreSQL: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while trying to ping PostgreSQL: %v", err)
	}

	return db, nil
}
