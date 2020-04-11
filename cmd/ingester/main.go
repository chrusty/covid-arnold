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
	logger.SetLevel(logrus.DebugLevel)

	config, err := configuration.Load()
	if err != nil {
		logger.WithError(err).Fatal("Unable to load config")
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
