package config


func LoadProdConfig() Config {
	return Config {
		ENV: "production",
		Port: GetEnv("PORT", "9999"),
	}
}