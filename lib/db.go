package lib

import (
	"database/sql"
	"log"
)

func PrepareDB(cfg *Config) *sql.DB {
	var db, err = sql.Open(cfg.DBType, "refthecord.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
