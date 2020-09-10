package server

// Config struct
type Config struct {
	BindAddr       string `toml:"bind_addr"`
	ShortURLLength int    `toml:"shorturl_length"`

	// database parameters
	MongodbConnection string `toml:"mongodb_conn"`
	DBName            string `toml:"dbname"`
	CollectionName    string `toml:"collection_name"`
}

// NewConfig - return new config
func NewConfig() *Config {
	return &Config{}
}
