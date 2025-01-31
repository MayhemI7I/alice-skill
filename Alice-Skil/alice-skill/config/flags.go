package config

import "github.com/spf13/pflag"

type Config struct {
	ServerAdress string 
	ServerPort   string
	BaseURL string
}

func InitConfig() *Config {
	cfg := &Config{}
	pflag.StringVarP(&cfg.ServerAdress, "server-address", "s", "localhost", "server address")
    pflag.StringVarP(&cfg.ServerPort, "server-port", "p", "8080", "server port")
	pflag.StringVarP(&cfg.BaseURL,"base-url", "b", "http://localhost:8080", "Base URL for return server")
    
	pflag.Parse()
	return cfg
}

