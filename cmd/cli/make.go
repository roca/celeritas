package main

import (
	"errors"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		err := doMakeMigrations(arg3)
		if err != nil {
			exitGracefully(err)
		}

	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}

	default:
		return errors.New("make requires a subcommand: (migration|model|handler)")
	}
	return nil
}
