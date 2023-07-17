package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		dbType := cel.DB.DataType
		if arg3 == "" {
			exitGracefully(errors.New("you must give the migration a name"))
		}

		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)

		upFile := cel.RooPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := cel.RooPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		err := copyFileFromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err!=nil {
			exitGracefully(err)
		}
		err = copyFileFromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err!=nil {
			exitGracefully(err)
		}

	default:
		return errors.New("make requires a subcommand: (migration|model|handler)")
	}
	return nil
}
