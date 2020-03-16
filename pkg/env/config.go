package env

// Config holds the env varibles that are used to configure the server
type Config struct {
	Port      string
	Host      string
	Dbport    int
	User      string
	Password  string
	Dbname    string
	Secretkey string
}
