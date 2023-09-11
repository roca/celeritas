package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func setup(arg1, agr2 string) {
	if arg1 != "new" && arg1 != "version" && arg1 != "help" {
		err := godotenv.Load()
		if err != nil {
			exitGracefully(err)
		}

		path, err := os.Getwd()
		if err != nil {
			exitGracefully(err)
		}

		cel.RooPath = path
		cel.DB.DataType = os.Getenv("DATABASE_TYPE")
	}
}

func getDSN() string {
	dbType := cel.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		}
		return dsn
	}
	return "mysql://" + cel.BuildDSN()
}

func showHelp() {
	color.Yellow(`Available commands:
	help	     		- show the help commands
	version	     		- print application version
	migrate	     		- run all up migrations that have not been run previously
	migrate down 		- reverses the most recent migration
	migrate reset 		- runs all down migrations in reverse order, and then runs all up migrations
	make migration <name>	- create two new up and down migrations in the migrations folder
	make auth 		- creates and runs migrations for the authentication tables, and creates models and middleware
	make handler <name>	- creates a stub handler in the handlers directory
	make model <name>	- creates a new model in the data directory
	make session 		- creates a table in the database as a session store
	make mail <name>	- creates two starter mail templates in the mail directory

	`)
}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	// check for an err before doing anything else
	if err != nil {
		return err
	}

	// check if current file is a directory
	if fi.IsDir() {
		return nil
	}

	// check if current file is a go file
	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}

	// we have a go file
	if matched {
		// read file contents
		read, err := os.ReadFile(path)
		if err != nil {
			exitGracefully(err)
		}

		// update file contents
		newContents := strings.Replace(string(read), "myapp", appURL, -1)

		// write the changed file
		err = os.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			exitGracefully(err)
		}
	}

	return nil
}

func updateSource() {
	// walk entire project folder, including sub folders
	err := filepath.Walk(".", updateSourceFiles)
	if err != nil {
		exitGracefully(err)
	}
}
