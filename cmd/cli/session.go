package main

func doSessionTable() error {

	err := makeMigrations("session")
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	return nil
}
