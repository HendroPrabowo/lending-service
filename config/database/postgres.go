package database

import (
	"context"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var Postgres *pg.DB

func InitPostgreOrm() {
	//host := os.Getenv("DATABASE_HOST")
	//port := os.Getenv("DATABASE_PORT")
	//user := os.Getenv("DATABASE_USER")
	//password := os.Getenv("DATABASE_PASSWORD")
	//database := os.Getenv("DATABASE_NAME")
	//
	//if host == "" || port == "" || user == "" || password == "" || database == "" {
	log.Info("connect to database using localhost")
	host := DATABASE_HOST
	port := DATABASE_PORT
	user := DATABASE_USER
	password := DATABASE_PASSWORD
	database := DATABASE_NAME
	//}

	log.Infof("HOST : %s, PORT : %s, USER : %s, PASS : %s, DATABASE : %s", host, port, user, password, database)

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
