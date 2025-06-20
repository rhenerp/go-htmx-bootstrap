package config

func LoadDevConfig() Config {
	return Config {
		ENV: "development",
		Port: "3000",
	}
}