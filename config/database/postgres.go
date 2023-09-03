package database

import (
	"context"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var Postgres *pg.DB

func InitPostgreOrm() {
	db := pg.Connect(&pg.Options{
		Addr:     HOST + ":" + PORT,
		User:     USER,
		Password: PASSWORD,
		Database: DATABASE,
	})
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	Postgres = db
	log.Info("connected to database POSTGRES")
}
