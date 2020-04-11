package storage

import "github.com/chrusty/covid-arnold/internal/models"

// StoreDailyRegions saves a list of DailyRegions to the DB
func (m *Manager) StoreDailyRegions(dailyRegions []*models.DailyRegion) error {

	// Insert the records
	for _, dailyRegion := range dailyRegions {
		m.dbConn.Create(dailyRegion)
	}

	return nil
}
