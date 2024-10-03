package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Online_Song_Libraries/config"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to database")
	return db, nil
}
