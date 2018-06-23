package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// We don't need to use lib/pq directly
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	reform "gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/milangrahovac/users/internal/web"
)

func main() {
	debug := os.Getenv("DEBUG")
	log := logrus.New()
	if len(debug) > 0 && debug != "0" {
		log.SetLevel(logrus.DebugLevel)
		log.Info("Debug mode is on.")
	}

	log.Info("Starting service...")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port is not set.")
	}

	db, err := startupDB()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: web.NewRouter(log, db),
	}

	log.Infof("Listening on the port %s", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func startupDB() (*reform.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if len(dbURL) == 0 {
		return nil, fmt.Errorf("the database is not set.")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Try ping to check connection
	attempts := 5
	for i := 0; i < attempts; i++ {
		err = conn.Ping()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("couldn't ping the DB: `%s`", err)
	}

	db := reform.NewDB(
		conn,
		postgresql.Dialect,
		reform.NewPrintfLogger(log.Printf),
	)
	return db, nil
}
