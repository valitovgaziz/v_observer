package apiserver

type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"`
}

// New conifg
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8282",
		LogLevel: "debug",
	}
}
