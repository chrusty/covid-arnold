package models

// DailyRegion represents a daily report for a specific region
type DailyRegion struct {
	Confirmed   int32
	CountryCode string
	CountryName string
	Date        string
	Deaths      int32
	Key         string
	Latitide    string
	Longitude   string
	Population  int64
	RegionCode  string
	RegionName  string
}
