package main


func doAuth() error {
	// migrations

	err:= doMakeMigrations("auth")
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copy some files
	err = copyFileFromTemplate("templates/data/user.go.txt", cel.RooPath + "/data/user.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/data/token.go.txt", cel.RooPath + "/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy over middleware
	err = copyFileFromTemplate("templates/middleware/auth.go.txt", cel.RooPath + "/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", cel.RooPath + "/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	return nil
}