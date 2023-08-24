package mailer

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var pool *dockertest.Pool
var resourse *dockertest.Resource
var mailer = Mail{
	Domain:      "localhost",
	Templates:   "./testdata/mail",
	Host:        "localhost",
	Port:        1026,
	Encryption:  "none",
	FromAddress: "me@here.com",
	FromName:    "Joe",
	Jobs:        make(chan Message, 1),
	Results:     make(chan Result, 1),
}

func setup() {
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	pool = p

	opts := dockertest.RunOptions{
		Repository:   "mailhog/mailhog",
		Tag:          "latest",
		Env:          []string{},
		ExposedPorts: []string{"1025", "8025"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"1025": {
				{HostIP: "0.0.0.0", HostPort: fmt.Sprintf("%d", mailer.Port)},
			},
			"8025": {
				{HostIP: "0.0.0.0", HostPort: "8026"},
			},
		},
	}

	resourse, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resourse)
		log.Fatalf("Could not start resource: %s", err)
	}

	time.Sleep(2 * time.Second)

	go mailer.ListenForMail()

	return
}

func teardown() {
	err := pool.Purge(resourse)
	if err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	return
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	// teardown()
	os.Exit(code)
}
