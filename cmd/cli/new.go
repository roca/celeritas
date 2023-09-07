package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

func doNew(appName string) {
	appName = strings.ToLower(appName)

	// sanitize the application name (convert url to single word)
	if strings.Contains(appName, "/") {
		// get the last part of the url
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[len(exploded)-1]
	}

	log.Println("App name is", appName)

	// git clone the skeleton application
	color.Green("\tCloning skeleton application")
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		// Auth: &http.BasicAuth{
		// 	Username: "abc123", // yes, this can be anything except an empty string
		// 	Password: os.Getenv("CELERITAS_SKELETON_ACCESS_TOKEN"),
		// },
		// Auth: &http.TokenAuth{
		// 	Token: os.Getenv("CELERITAS_SKELETON_ACCESS_TOKEN"),
		// },
		// URL:      "git@github.com:roca/celeritas-skeleton.git",
		URL:      "https://github.com/roca/celeritas-skeleton.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		exitGracefully(err)
	}

	// remove .git directory
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		exitGracefully(err)
	}

	// create a ready to go .env file

	// create a makefile

	// update the go.mod file

	// update existing .go files with correct name/imports

	// run go mod tidy in the project directory
}
