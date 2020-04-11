package main

import (
	"time"

	"github.com/chrusty/covid-arnold/internal/configuration"
	"github.com/chrusty/covid-arnold/internal/ingester"
	"github.com/chrusty/covid-arnold/internal/storage"

	"github.com/sirupsen/logrus"
)

const (
	datasetURL = "https://open-covid-19.github.io/data/data.json"
)

func main() {

	// Dependencies
	logger := logrus.New()

	// Load config (from env-vars)
	config, err := configuration.Load()
	if err != nil {
		logger.WithError(err).Fatal("Unable to load config")
	}

	// Set the log-level (if it parses)
	if logrusLevel, err := logrus.ParseLevel(config.Logging.Level); err != nil {
		logger.WithError(err).Warnf("Configured log-level (%s) is invalid", config.Logging.Level)
	} else {
		logger.SetLevel(logrusLevel)
	}

	// Prepare a storage managerL
	storageManager, err := storage.New(logger, config.Database.ConnectionString())
	if err != nil {
		logger.WithError(err).Fatal("Unable to prepare a storage.Manager")
	}

	// Start the ingester
	ingester := ingester.New(datasetURL, logger, storageManager)
	ingester.Start(time.Hour)
}
