package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"

	"github.com/mahiro72/go-api-template/pkg/env"
)

func Init() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", env.CONFIG.MYSQL.DSN)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		spentSec := 0
		for err != nil && spentSec <= 60 {
			db, err = sqlx.Connect("mysql", env.CONFIG.MYSQL.DSN)
			log.Printf("failed: connect to db, time: %ds ...\n", spentSec)
			spentSec += 5
			time.Sleep(time.Second * 5)
		}
	}

	log.Println("success: connect to db")
	return db, nil
}
