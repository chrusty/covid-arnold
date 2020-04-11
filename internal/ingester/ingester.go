package ingester

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chrusty/covid-arnold/internal/models"
	"github.com/chrusty/covid-arnold/internal/storage"
	"github.com/sirupsen/logrus"
)

// Ingester is a self-contained data ingester
type Ingester struct {
	dataURL        string
	logger         *logrus.Logger
	storageManager *storage.Manager
}

// New returns a configured Ingester
func New(dataURL string, logger *logrus.Logger, storageManager *storage.Manager) *Ingester {
	return &Ingester{
		dataURL:        dataURL,
		logger:         logger,
		storageManager: storageManager,
	}
}

// Start tells the ingester to begin its ingestion cycle
func (i *Ingester) Start(frequency time.Duration) {
	i.logger.WithField("frequency", frequency).Info("Ingester starting")
	for {
		i.ingest()
		time.Sleep(frequency)
	}
}

// ingest actually downloads and inserts the data
func (i *Ingester) ingest() error {
	var dailyRegions []*models.DailyRegion

	i.logger.WithField("dataURL", i.dataURL).Info("Ingesting")

	// Get the data
	resp, err := http.Get(i.dataURL)
	if err != nil {
		i.logger.WithError(err).Error("Unable to get data")
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&dailyRegions); err != nil {
		i.logger.WithError(err).Error("Unable to decode response")
	}
	i.logger.WithField("records", len(dailyRegions)).Debug("Downloaded data")

	return i.storageManager.StoreDailyRegions(dailyRegions)
}
