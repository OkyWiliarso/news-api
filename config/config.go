package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "qwerty123",
			Host:     "localhost",
			Port:     "3306",
			Name:     "news_db",
		},
	}
}
