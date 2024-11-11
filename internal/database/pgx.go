package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jordanmarcelino/terradiscover-backend/internal/config"
	"github.com/jordanmarcelino/terradiscover-backend/internal/logger"
)

func InitStdLib(cfg *config.Config) *sql.DB {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DbName,
		dbCfg.Port,
		dbCfg.Sslmode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Log.Fatalf("error initializing database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Log.Fatalf("error connecting to database: %v", err)
	}

	db.SetMaxIdleConns(dbCfg.MaxIdleConn)
	db.SetMaxOpenConns(dbCfg.MaxOpenConn)
	db.SetConnMaxLifetime(time.Duration(dbCfg.MaxConnLifetimeMinute) * time.Minute)

	return db
}