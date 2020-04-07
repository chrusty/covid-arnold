package main

import (
	"time"

	"github.com/chrusty/covid-arnold/internal/ingester"
	"github.com/sirupsen/logrus"
)

const (
	datasetURL = "https://open-covid-19.github.io/data/data.json"
)

func main() {

	// Dependencies
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	ingester := ingester.New(datasetURL, logger)

	// Start the ingester
	ingester.Start(time.Hour)
}
