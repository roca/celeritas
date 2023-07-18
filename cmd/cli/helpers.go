package main

import (
	"os"

	"github.com/joho/godotenv"
)

func setup() {
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

func getDSN() string {
	dbType := cel.DB.DataType

	if dbType == "pgx" {
		dbType = "postgress"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {

		} else {

		}

	} else {

	}
}
