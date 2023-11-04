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
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	database := os.Getenv("DATABASE_NAME")

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		log.Info("POSTGRES : config from localhost")
		host = DATABASE_HOST
		port = DATABASE_PORT
		user = DATABASE_USER
		password = DATABASE_PASSWORD
		database = DATABASE_NAME
	}

	opt := &pg.Options{
		Addr:     host + ":" + port,
		User:     user,
		Password: password,
		Database: database,
	}

	var err error
	pgConnectionString := os.Getenv("DATABASE_CONNECTION_STRING")
	if pgConnectionString != "" {
		log.Infof("POSTGRES : connect using connection string url")
		opt, err = pg.ParseURL(pgConnectionString)
		if err != nil {
			log.Fatal(err)
		}
	}

	db := pg.Connect(opt)
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	Postgres = db
	log.Info("POSTGRES : connected to database POSTGRES")
}
