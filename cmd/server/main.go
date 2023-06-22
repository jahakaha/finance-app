package main

import (
	"finance-app/internal/api"
	"finance-app/internal/config"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Create Server object and start listener
func main() {

	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("version", config.Version).Debug("")

	router, err := api.NewRouter()
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	const addr = "0.0.0.0:8080"

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Debug("Server failed")
	}

}
