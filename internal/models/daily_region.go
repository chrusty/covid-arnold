package models

// DailyRegion represents a daily report for a specific region
type DailyRegion struct {
	Confirmed   float32
	CountryCode string `gorm:"PRIMARY_KEY"`
	CountryName string `gorm:"index:country_name"`
	Date        string `gorm:"type:date;PRIMARY_KEY"`
	Deaths      float32
	Key         string `gorm:"PRIMARY_KEY"`
	Latitude    string
	Longitude   string
	Population  float32
	RegionCode  string `gorm:"index:region_code"`
	RegionName  string `gorm:"index:region_name"`
}
