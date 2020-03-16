package env

// DBConfig holds the env variables that are used to configure the DB
type DBConfig struct {
	Host     string
	Dbport   int // The port at which the DB is running
	User     string
	Password string
	Dbname   string
}

// Config holds the env varibles that are used to configure the server
type Config struct {
	Port      string // The port at which the server runs
	Secretkey string // The secret key for decoding JWT token
	DB        *DBConfig
}

// LoadConfig creates and returns a struct with config values
func LoadConfig() *Config {
	// TODO: Load configuration values from an external env file that is not checked into version control
	// These hardcoded values are only for testing during development
	return &Config{
		Port:      ":8080",
		Secretkey: "supersecretkey",
		DB: &DBConfig{
			Host:     "db",
			Dbport:   5432,
			User:     "admin",
			Password: "password",
			Dbname:   "db",
		},
	}
}
