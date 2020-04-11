package configuration

import "fmt"

// Database configures a Postgres DB
type Database struct {
	DBName   string `env:"DB_SCHEMA" envDefault:"covid"`
	Hostname string `env:"DB_HOSTNAME" envDefault:"localhost"`
	Password string `env:"DB_PASSWORD" envDefault:"arnold"`
	Username string `env:"DB_USER_NAME" envDefault:"covid"`
}

// ConnectionString returns a Postgres connection string based on the configured values
func (d *Database) ConnectionString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", d.Hostname, d.Username, d.DBName, d.Password)
}
