package db

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // driver for open postgres connection

	"github.com/trantrongkim98/example-go/cmd/migrator/config"
	"github.com/trantrongkim98/example-go/domain"
)

// PGConnector store implement open for postgres
type PGConnector struct{}

// NewPGConnector create new PGConnector
func NewPGConnector() *PGConnector { return &PGConnector{} }

// Open open new connection to posgres using config
func (c *PGConnector) Open(cfg *config.Config) (*sql.DB, error) {
	sslmode := "disable"
	if cfg.DBSSLModeOption == "enable" {
		sslmode = "require"
	}

	dbstring := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=%s",
		cfg.DBUserName,
		cfg.DBName,
		sslmode,
		cfg.DBPassword,
		cfg.DBHostname,
		cfg.DBPort,
	)

	return sql.Open("postgres", dbstring)
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
