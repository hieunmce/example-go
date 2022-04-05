package db

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // driver for open postgres connection

	"example.com/m/cmd/migrator/config"
	"example.com/m/domain"
)

// PGConnector store implement open for postgres
type PGConnector struct{}

// NewPGConnector create new PGConnector
func NewPGConnector() *PGConnector { return &PGConnector{} }

// Open open new connection to posgres using config
func (c *PGConnector) Open(cfg *config.Config) (*sql.DB, error) {
	sslMode := "disable"
	if cfg.DBSSLModeOption == "enable" {
		sslMode = "require"
	}

	dbString := fmt.Sprintf("user=%s dbname=%s sslMode=%s password=%s host=%s port=%s",
		cfg.DBUserName,
		cfg.DBName,
		sslMode,
		cfg.DBPassword,
		cfg.DBHostname,
		cfg.DBPort,
	)

	return sql.Open("postgres", dbString)
}

// InitModel .
func (c PGConnector) InitModel(cfg *config.Config) error {

	connectionString := fmt.Sprintf("host = %s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHostname,
		cfg.DBPort,
		cfg.DBUserName,
		cfg.DBPassword,
		cfg.DBName,
	)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	return db.AutoMigrate(
		&domain.User{},
	).Error
}
