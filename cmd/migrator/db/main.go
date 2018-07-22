package db

import (
	"database/sql"
	"errors"

	"github.com/hieunmce/example-go/cmd/migrator/config"
)

const (
	dbTypePostgres = "postgres"
)

var (

	// ErrWrongDBType error for using wrong db type in config
	ErrWrongDBType = errors.New("Config using wrong or not support Database type")
)

// Connector interface for open a db connection with config
type Connector interface {
	Open(config.Config) (*sql.DB, error)
}

// NewConnection open new db connection using config
func NewConnection(cfg *config.Config) (*sql.DB, error) {
	if cfg.DBType == dbTypePostgres {
		return NewPGConnector().Open(cfg)
	}

	return nil, ErrWrongDBType
}

// ModelInitiator init model by db
type ModelInitiator interface {
	InitModel(cfg *config.Config) error
}

// InitModel from config
func InitModel(cfg *config.Config) error {
	if cfg.DBType == dbTypePostgres {
		return NewPGConnector().InitModel(cfg)
	}

	return ErrWrongDBType
}
