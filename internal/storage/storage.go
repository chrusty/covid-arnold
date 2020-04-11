package storage

import (
	"github.com/chrusty/covid-arnold/internal/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres dialect for Gorm
	"github.com/sirupsen/logrus"
)

// Manager has everything we need to interact with the database
type Manager struct {
	dbConn *gorm.DB
}

// New returns a fully configured storage interface
func New(logger *logrus.Logger, connectionString string) (*Manager, error) {

	// Prepare a DB connection
	dbConn, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Migrate tables for our models
	dbConn.AutoMigrate(&models.DailyRegion{})

	return &Manager{
		dbConn: dbConn,
	}, nil
}
