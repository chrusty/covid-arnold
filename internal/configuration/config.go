package configuration

import (
	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

// Config contains config structs for each part of the app
type Config struct {
	Database Database
	Logging  Logging
}

// Load reads in supplied environment vars
func Load() (*Config, error) {
	cfg := &Config{}

	for _, c := range []interface{}{
		cfg,
		&cfg.Database,
		&cfg.Logging,
	} {
		if err := env.Parse(c); err != nil {
			return nil, errors.Wrap(err, "env.Parse(..) failed to load configuration")
		}
	}

	return cfg, nil
}
