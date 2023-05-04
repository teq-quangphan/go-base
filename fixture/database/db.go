package database

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-base/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// list tables in DB.
var tables = []string{
	"users",
}

type Database struct {
	DB *gorm.DB
}

func InitDatabse() *Database {
	conf.SetEnv()
	cfg := conf.GetConfig()
	connectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		cfg.MySQL.DBTestUser,
		cfg.MySQL.DBTestPass,
		cfg.MySQL.DBTestHost,
		cfg.MySQL.DBTestPort,
		cfg.MySQL.DBTestName,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: connectionString}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return &Database{DB: db.Session(&gorm.Session{})}
}
func (d *Database) TruncateTables() {
	for i := range tables {
		err := d.DB.Table(tables[i]).Exec(fmt.Sprintf("TRUNCATE TABLE %s", tables[i])).Error
		if err != nil {
			logrus.Fatal(err.Error())
		}
	}
}
func (d *Database) ExecFixture(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	for scanner.Scan() {
		query := scanner.Text()
		if err = d.DB.Exec(query).Error; err != nil {
			return err
		}
	}

	return nil
}
func (d *Database) GetClient() *gorm.DB {
	return d.DB
}
