package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/madeindra/devoria-workshop-to-challenge/domain/account"
	"github.com/madeindra/devoria-workshop-to-challenge/domain/article"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/bcrypt"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/config"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/constant"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/jwt"
)

func main() {
	// config init
	cfg := config.New()

	// gorm database init
	db, err := sql.Open("mysql", cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.Database.MaxOpenConnections)

	// dependencies init
	validator := validator.New()
	router := mux.NewRouter()
	bcrypt := bcrypt.NewBcrypt(cfg.Bcrypt.HashCost)
	jsonWebToken := jwt.NewJSONWebToken(cfg.Jwt.PrivateKey, cfg.Jwt.PublicKey)

	// repo, usecase
	accountRepo := account.NewAccountRepository(db, constant.TableAccount)
	accountUsecase := account.NewAccountUsecase(accountRepo, bcrypt, jsonWebToken)

	articleRepo := article.NewAccountRepository(db, constant.TableArticle)
	articleUsecase := article.NewArticleUsecase(articleRepo, accountRepo)

	// router to handler mapping
	account.NewAccountHandler(router, validator, accountUsecase)
	article.NewArticleHandler(router, validator, articleUsecase)

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

	db.Close()
}
