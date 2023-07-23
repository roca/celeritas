package main

import (
	"errors"
	"fmt"
	"time"
)

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	// run the migration command
	switch arg2 {
	case "up":
		err := cel.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := cel.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := cel.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := cel.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = cel.MigrateUp(dsn)
		if err != nil {
			return err
		}

	default:
		showHelp()
	}

	return nil
}

func doMakeMigrations(modelName string) error {
	var migrationFilePrefix  string
	dbType := cel.DB.DataType
	if modelName == "" {
		return errors.New("you must give the migration a name")
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), modelName)

	upFile := cel.RooPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := cel.RooPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	if modelName == "auth" {
		migrationFilePrefix = fmt.Sprintf("templates/migrations/auth_tables.%s", dbType)
	} else {
		migrationFilePrefix = fmt.Sprintf("templates/migrations/migration.%s", dbType)
	}

	err := copyFileFromTemplate(migrationFilePrefix+".up.sql", upFile)
	if err != nil {
		return err
	}
	err = copyFileFromTemplate(migrationFilePrefix+".down.sql", downFile)
	if err != nil {
		return err
	}
	return nil
}
