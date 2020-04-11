package configuration

// Logging configures the logger
type Logging struct {
	Level string `env:"LOGGING_LEVEL" envDefault:"INFO"`
}
