package storage

import "github.com/chrusty/covid-arnold/internal/models"

// StoreDailyRegions saves a list of DailyRegions to the DB
func (m *Manager) StoreDailyRegions(dailyRegions []*models.DailyRegion) error {

	// Insert the records
	for _, dailyRegion := range dailyRegions {
		if err := m.dbConn.Create(dailyRegion).Error; err != nil {
			m.logger.WithError(err).Warn("Error inserting a row")
		}
	}

	return nil
}
