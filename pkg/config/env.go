package config

import "time"

var ENV = struct {
	// MasterKeys is an array of keys for encryption and decryption
	// Keys are recommended to have 32 random bytes
	// Decryption: try to use all the keys
	// Encryption: use the first key
	MasterKeys []string `env:"DSRHUB_MASTER_KEYS" envDefault:"please_change_me!!" envSeparator:","`
	LoggerEnv  string   `env:"DSRHUB_LOGGER_ENV" envDefault:"production"`

	AdminServerHostPort string `env:"DSRHUB_ADMIN_SERVER_HOST_PORT" envDefault:"localhost:28001"`
	ApiServerHostPort   string `env:"DSRHUB_API_SERVER_HOST_PORT" envDefault:"localhost:28000"`

	DBDriver                  string        `env:"DSRHUB_DB_DRIVER" envDefault:"sqlite3"`
	DBConnectionStr           string        `env:"DSRHUB_DB_CONNECTION_STR" envDefault:":memory:"`
	DBConnectionRetryAttempts uint          `env:"DSRHUB_DB_CONNECTION_RETRY_ATTEMPTS" envDefault:"9"`
	DBConnectionRetryDelay    time.Duration `env:"DSRHUB_DB_CONNECTION_RETRY_DELAY" envDefault:"100ms"`
	DBConnectionDebug         bool          `env:"DSRHUB_DB_CONNECTION_DEBUG" envDefault:"false"`

	DSLPath string `env:"DSRHUB_DSL_PATH" envDefault:"./demo_dsl.yaml"`
}{}
