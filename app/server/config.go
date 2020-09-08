package server

// Config struct
type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// NewConfig - return new config
func NewConfig() *Config {
	return &Config{}
}
