package storage

import (
	"github.com/chrusty/covid-arnold/internal/models"
)

// StoreDailyRegions saves a list of DailyRegions to the DB
func (m *Manager) StoreDailyRegions(dailyRegions []*models.DailyRegion) error {
	var rowsInserted int64

	// Insert the records
	for _, dailyRegion := range dailyRegions {
		response := m.dbConn.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(dailyRegion)
		if response.Error != nil {
			m.logger.WithError(response.Error).Warn("Error inserting a row")
		}
		rowsInserted += response.RowsAffected
	}

	m.logger.WithField("records_inserted", rowsInserted).Debug("Finished storing daily-regions records")

	return nil
}

// StoreDailyRegionsBlob saves a list of DailyRegions to the DB
func (m *Manager) StoreDailyRegionsBlob(dailyRegions []*models.DailyRegion) error {

	// Insert the records
	response := m.dbConn.Create(dailyRegions)
	if response.Error != nil {
		m.logger.WithError(response.Error).WithField("records_inserted", response.RowsAffected).Warn("Error inserting rows")
		return response.Error
	}

	m.logger.WithField("records_inserted", response.RowsAffected).Debug("Finished storing daily-regions records")

	return nil
}
