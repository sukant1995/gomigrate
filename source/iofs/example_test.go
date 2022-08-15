//go:build go1.16
// +build go1.16

package iofs_test

import (
	"embed"
	"log"

	"github.com/sukant1995/gomigrate"
	_ "github.com/sukant1995/gomigrate/database/postgres"
	"github.com/sukant1995/gomigrate/source/iofs"
)

//go:embed testdata/migrations/*.sql
var fs embed.FS

func Example() {
	d, err := iofs.New(fs, "testdata/migrations")
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, "postgres://postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		// ...
	}
	// ...
}
