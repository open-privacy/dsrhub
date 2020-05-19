package config

var ENV = struct {
	// MasterKeys is an array of keys for encryption and decryption
	// Keys are recommended to have 32 random bytes
	// Decryption: try to use all the keys
	// Encryption: use the first key
	MasterKeys []string `env:"DSRHUB_MASTER_KEYS" envDefault:"" envSeparator:","`
	LoggerEnv  string   `env:"DSRHUB_LOGGER_ENV" envDefault:"production"`
}{}
