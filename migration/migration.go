package migration

import (
	"log"

	"go-base/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) {
	getDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(getDB, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migration", conf.GetConfig().MySQL.DBName, driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println("Up done!")
}
