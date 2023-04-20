package main

import (
	"fmt"
	"go-base/conf"
	serviceHttp "go-base/delivery/http"
	"go-base/migration"
	"go-base/repository"
	"go-base/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

const VERSION = "1.0.0"

// @title Example API
// @version 1.0

// @BasePath /api
// @schemes http http

// @securityDefinitions.apikey AuthToken
// @in header
// @name Authorization

// @description Transaction API.
func main() {
	conf.SetEnv()

	confMysql := conf.GetConfig().MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", confMysql.DBUser, confMysql.DBPass, confMysql.DBHost, confMysql.DBPort, confMysql.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.New(db)
	uc := usecase.New(repo)

	//migration
	migration.Up(db)

	h := serviceHttp.NewHTTPHandler(uc)
	//go func() {
	//	h.Listener = httpL
	//	errs <- h.Start("")
	//}()
	if err := h.Start("localhost:8088"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
