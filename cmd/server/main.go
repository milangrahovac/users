package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/milangrahovac/users/internal/web"
)

func main() {
	debug := os.Getenv("DEBUG")
	log := logrus.New()

	if len(debug) > 0 && debug != "0" {
		log.SetLevel(logrus.DebugLevel)
		log.Info("Debug mode is on")
	}
	log.Info("Starting service....")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port is not set.")
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: web.NewRouter(log),
	}

	log.Infof("Listening on the port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
