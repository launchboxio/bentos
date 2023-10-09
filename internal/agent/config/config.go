package config

type Config struct {
}

func Load(filePath string) (*Config, error) {
	return &Config{}, nil
}
