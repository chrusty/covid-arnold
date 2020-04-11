package configuration

import "fmt"

// Database configures a Postgres DB
type Database struct {
	Hostname string `env:"DB_PASSWORD" envDefault:"localhost"`
	Password string `env:"DB_PASSWORD" envDefault:"changeme"`
	Schema   string `env:"DB_SCHEMA" envDefault:"postgres"`
	Username string `env:"DB_USER_NAME" envDefault:"postgres"`
}

// ConnectionString returns a Postgres connection string based on the configured values
func (d *Database) ConnectionString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", d.Hostname, d.Username, d.Schema, d.Password)
}
