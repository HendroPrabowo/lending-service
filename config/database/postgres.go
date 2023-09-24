package database

import (
	"context"
	"os"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var Postgres *pg.DB

func InitPostgreOrm() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		log.Info("connect to database using localhost")
		host = HOST
		port = PORT
		user = USER
		password = PASSWORD
		database = DATABASE
	}

	db := pg.Connect(&pg.Options{
		Addr:     host + ":" + port,
		User:     user,
		Password: password,
		Database: database,
	})
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	Postgres = db
	log.Info("connected to database POSTGRES")
}
