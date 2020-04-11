package models

import (
	"github.com/jinzhu/gorm"
)

// DailyRegion represents a daily report for a specific region
type DailyRegion struct {
	gorm.Model
	Confirmed   int32
	CountryCode string `gorm:"PRIMARY_KEY"`
	CountryName string `gorm:"index:country_name"`
	Date        string `gorm:"PRIMARY_KEY"`
	Deaths      int32
	Key         string
	Latitide    string
	Longitude   string
	Population  int64
	RegionCode  string `gorm:"PRIMARY_KEY,NULL"`
	RegionName  string `gorm:"index:region_name"`
}
