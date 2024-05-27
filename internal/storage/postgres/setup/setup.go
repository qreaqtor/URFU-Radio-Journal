package setupst

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func GetConnect(user, password, host, dbName string, port, tryConn int) (*sql.DB, error) {
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

	timer := time.NewTicker(time.Second)
	for i := 1; i <= tryConn; i++ {
		<-timer.C
		err = db.Ping()
		if err == nil {
			break
		}
		fmt.Printf("attempt %d to ping PostgreSQL: %v\n", i, err)
	}
	timer.Stop()
	if err != nil {
		return nil, fmt.Errorf("error while trying to ping PostgreSQL: %v", err)
	}
	// err = db.Ping()
	// if err != nil {
	// 	return nil, fmt.Errorf("error while trying to ping PostgreSQL: %v", err)
	// }

	fmt.Println("Success connection to PostgreSQL!")
	return db, nil
}
