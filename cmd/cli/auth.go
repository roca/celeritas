package main

import "github.com/fatih/color"

func doAuth() error {
	// migrations

	err := doMakeMigrations("auth")
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copy some files
	err = copyFileFromTemplate("templates/data/user.go.txt", cel.RooPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/data/token.go.txt", cel.RooPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy over middleware
	err = copyFileFromTemplate("templates/middleware/auth.go.txt", cel.RooPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", cel.RooPath+"/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow(" - users, tokens, and remember_tokens migrations created and executed")
	color.Yellow(" - user and token models created")
	color.Yellow(" - auth middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token models in data/models.go, and to add appropriate middleware to your routes!")

	return nil
}
