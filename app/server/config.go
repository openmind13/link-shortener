package server

// Config struct
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	MongoDBConn string `toml:"mongodb_conn"`
}

// NewConfig - return new config
func NewConfig() *Config {
	return &Config{}
}
