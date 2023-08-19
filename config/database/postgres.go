package database

import (
	"context"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var Postgres *pg.DB

func InitPostgreOrm() {
	opt, err := pg.ParseURL(PostgreDev)
	if err != nil {
		log.Fatal(err)
	}
	db := pg.Connect(opt)
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	Postgres = db
	log.Info("Connected to database POSTGRES")
}
