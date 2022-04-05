package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose"

	// Init DB drivers.
	"example.com/m/cmd/migrator/config"
	dbconn "example.com/m/cmd/migrator/db"
)

var (
	flags = flag.NewFlagSet("migrator", flag.ExitOnError)
	dir   = flags.String("dir", "cmd/migrator/migration", "directory with migration files")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	args := flags.Args()

	// create migrations
	if len(args) > 1 && args[0] == "create" {
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("migrator run: %v", err)
		}
		return
	}

	if len(args) < 1 {
		flags.Usage()
		return
	}

	if args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	cfg, err := config.GetBy(config.NewFileReader(".env_migrator", "."))
	if err != nil {
		log.Fatalf("can't read config by error: %v", err)
	}

	db, err := dbconn.NewConnection(cfg)
	if err != nil {
		log.Fatalf("can't connect to db by error: %v", err)
	}

	command := args[0]
	var arguments []string
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	// init model
	if args[0] == "init" {
		if err = dbconn.InitModel(cfg); err != nil {
			log.Fatalf("Failed to create tables by models by error: %v", err)
		}
		return
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("migrator run: %v", err)
	}
}

func usage() {
	fmt.Print(usagePrefix)
	flags.PrintDefaults()
	fmt.Print(usageCommands)
}

var (
	usagePrefix = `Usage: migrator [OPTIONS] COMMAND
Examples:
    migrator init
    migrator status
    migrator create add_some_column sql
    migrator create fetch_user_data go
    migrator up
Options:
`

	usageCommands = `
Commands:
    init                 Create tables base by models
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version
`
)
