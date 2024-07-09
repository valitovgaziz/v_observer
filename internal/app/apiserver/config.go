package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// New conifg
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8282",
		LogLevel: "debug",
	}
}
