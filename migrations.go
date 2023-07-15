package celeritas

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
)

// MigrateUp runs the migrations up
func (c *Celeritas) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+c.RooPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil {
		log.Println("Error running migration:", err)
		return err
	}

	return nil
}
