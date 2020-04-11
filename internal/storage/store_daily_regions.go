package storage

import (
	"github.com/chrusty/covid-arnold/internal/models"
)

// StoreDailyRegions saves a list of DailyRegions to the DB
func (m *Manager) StoreDailyRegions(dailyRegions []*models.DailyRegion) error {
	var rowsInserted int32

	// Insert the records
	for _, dailyRegion := range dailyRegions {
		if err := m.dbConn.Create(dailyRegion).Error; err != nil {
			m.logger.WithError(err).Warn("Error inserting a row")
		} else {
			rowsInserted++
		}
	}

	m.logger.WithField("records_inserted", rowsInserted).Debug("Finished storing daily-regions records")

	return nil
}
