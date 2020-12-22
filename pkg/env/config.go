package env

import (
	"strconv"

	"github.com/joho/godotenv"
)

// DBConfig holds the env variables that are used to configure the DB
type DBConfig struct {
	Host     string
	Dbport   int // The port at which the DB is running
	User     string
	Password string
	Dbname   string
	SSLMode  string
}

// Config holds the env varibles that are used to configure the server
type Config struct {
	Port      string // The port at which the server runs
	Secretkey string // The secret key for decoding JWT token
	LastLevel int    // The last level of the game
	DB        *DBConfig
}

// LoadConfig creates and returns a struct with config values
func LoadConfig() (*Config, error) {
	env, err := godotenv.Read()
	if err != nil {
		return nil, err
	}
	dbPort, _ := strconv.Atoi(env["DB_PORT"])
	ll, _ := strconv.Atoi(env["LAST_LEVEL"])
	return &Config{
		Port:      ":" + env["PORT"],
		Secretkey: env["SECRET_KEY"],
		LastLevel: ll,
		DB: &DBConfig{
			Host:     env["DB_HOST"],
			Dbport:   dbPort,
			User:     env["DB_USER"],
			Password: env["DB_PASSWORD"],
			Dbname:   env["DB_NAME"],
			SSLMode:  env["SSLMODE"],
		},
	}, nil
}
