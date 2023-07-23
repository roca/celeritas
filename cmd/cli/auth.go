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

	return nil
}