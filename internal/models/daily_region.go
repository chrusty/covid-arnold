package models

// DailyRegion represents a daily report for a specific region
type DailyRegion struct {
	Confirmed   float32
	CountryCode string `gorm:"PRIMARY_KEY"`
	CountryName string `gorm:"index:country_name"`
	Date        string `gorm:"PRIMARY_KEY"`
	Deaths      float32
	Key         string
	Latitude    string
	Longitude   string
	Population  float32
	RegionCode  string `gorm:"PRIMARY_KEY,NULL"`
	RegionName  string `gorm:"index:region_name"`
}
