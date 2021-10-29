package cachehero

// Config is used by newConn to get a new connection to a cache client
type Config struct {
	Driver   string `json:"driver,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     uint   `json:"port,omitempty"`
	Database string `json:"database,omitempty"`
	Username string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

// NewConfig obtains a new Config model
func NewConfig(driver string, host string, port uint, database string, user string, password string) Config {
	return Config{Driver: driver, Host: host, Port: port, Database: database, Username: user, Password: password}
}
