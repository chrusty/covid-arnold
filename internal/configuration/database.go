package configuration

import "fmt"

// Database configures a Postgres DB
type Database struct {
	Hostname string `env:"DB_HOSTNAME" envDefault:"localhost"`
	Name     string `env:"DB_NAME" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Username string `env:"DB_USER_NAME" envDefault:"postgres"`
}

// ConnectionString returns a Postgres connection string based on the configured values
func (d *Database) ConnectionString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", d.Hostname, d.Username, d.Name, d.Password)
}
