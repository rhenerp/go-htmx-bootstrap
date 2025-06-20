package config


func LoadStagingConfig() Config {
	return Config {
		ENV: "staging",
		Port: GetEnv("PORT", "3003"),
	}
}