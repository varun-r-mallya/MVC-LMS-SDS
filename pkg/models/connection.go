package models

import (
	"database/sql"

	"context"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
)

func Connection() (*sql.DB, error) {

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		neem.Critial(err, "Database connection failed")

	}
	neem.Log("Connected to DB Successfully")
	return db, err
}
