package main


func doAuth() error {
	// migrations

	err:= doMakeMigrations("users")
	if err != nil {
		return err
	}

	// run migrations

	// copy some files

	return nil
}