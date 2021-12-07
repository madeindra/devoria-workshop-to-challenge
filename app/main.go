package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/madeindra/devoria-workshop-to-challenge/domain/account"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// config init
	//cfg := config.New()

	// gorm init
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// validator init
	validator := validator.New()

	// router init
	router := mux.NewRouter()

	// repo, usecase, hadnler init
	accountRepo := account.NewAccountRepository(db, "account")
	accountUsecase := account.NewAccountUsecase(accountRepo)
	account.NewAccountHandler(router, validator, accountUsecase)

	// server init
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: router,
	}

	// graceful shutdown
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT)
	<-sigterm

	fmt.Println("shutting down application ...")

	server.Shutdown(context.Background())

	sqlDB.Close()
}
