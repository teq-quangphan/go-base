package main

import (
	"fmt"
	"go-base/conf"
	grpc2 "go-base/delivery/grpc"
	serviceHttp "go-base/delivery/http"
	"go-base/proto"
	"go-base/repository"
	"go-base/usecase"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
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
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", confMysql.DBUser, confMysql.DBPass, confMysql.DBHost, confMysql.DBPort, confMysql.DBName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", confMysql.DBUser, confMysql.DBPass, confMysql.DBHost, confMysql.DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.New(db)
	uc := usecase.New(repo)
	go RunGRPC(uc)

	//migration
	//migration.Up(db)

	h := serviceHttp.NewHTTPHandler(uc)
	//go func() {
	//	h.Listener = httpL
	//	errs <- h.Start("")
	//}()
	if err := h.Start("0.0.0.0:8088"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func RunGRPC(uc *usecase.UseCase) {
	port := ":8888"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &grpc2.ServerGRPC{UseCase: uc})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
