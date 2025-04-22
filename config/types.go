package config

type Config struct {
	Server ServerConfig `json:"server"`
}

type ServerConfig struct {
	Port uint   `json:"port"`
	Host string `json:"host"`
}
