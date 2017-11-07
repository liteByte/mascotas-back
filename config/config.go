package config

type Config struct {
	ENV         string
	PORT        string
	JWT_SECRET  string
	DB_TYPE     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		config := newConfigLocal()
		instance = &config
	}
	return instance
}

func newConfigLocal() Config {
	return Config{
		ENV:         "develop",
		PORT:        "8083",
		JWT_SECRET:  "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
		DB_TYPE:     "mysql",
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_HOST:     "localhost",
		DB_PORT:     "3306",
		DB_NAME:     "mascotas-db",
	}
}
