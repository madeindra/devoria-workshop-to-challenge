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
	_ "github.com/joho/godotenv/autoload"
	"github.com/madeindra/devoria-workshop-to-challenge/domain/account"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/bcrypt"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/config"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/constant"
	"gorm.io/gorm"
)

func main() {
	// config init
	cfg := config.New()

	// gorm database init
	db, err := gorm.Open(cfg.Gorm.DB)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(cfg.Gorm.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.Gorm.MaxOpenConnections)

	// dependencies init
	validator := validator.New()
	router := mux.NewRouter()
	bcrypt := bcrypt.NewBcrypt(cfg.Bcrypt.HashCost)

	// repo, usecase, hadnler init
	accountRepo := account.NewAccountRepository(db, constant.TableAccount)
	accountUsecase := account.NewAccountUsecase(accountRepo, bcrypt)
	account.NewAccountHandler(router, validator, accountUsecase)

	// server init
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.App.Port),
		Handler: router,
	}

	// graceful shutdown
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT)
	<-sigterm

	fmt.Println(constant.MessageGracefulShutown)

	server.Shutdown(context.Background())

	sqlDB.Close()
}
