package postgrest

import (
	"database/sql"
	"fmt"
	"time"
	"urfu-radio-journal/internal/config"

	_ "github.com/lib/pq"
)

func GetConnect(conf config.PostgresConfig, ssl bool) (*sql.DB, error) {
	sslMode := "disable"
	if ssl {
		sslMode = "enable"
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Database,
		sslMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to PostgreSQL: %v", err)
	}

	tiker := time.NewTicker(time.Second)
	for i := 0; i < conf.ConnAttempts; i++ {
		_, ok := <-tiker.C
		if !ok {
			break
		}
		err = db.Ping()
		if err == nil {
			tiker.Stop()
			return db, nil
		}
	}

	return nil, fmt.Errorf("can't connect to PostgreSQL")
}
