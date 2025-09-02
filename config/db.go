package config

type DbConfig struct {
	Host     string `yaml:"HOST" env:"DB_HOST"`
	Username string `yaml:"USERNAME" env:"DB_USERNAME"`
	Password string `yaml:"PASSWORD" env:"DB_PASSWORD"`
	Port     string `yaml:"PORT" env:"DB_PORT"`
	DBName   string `yaml:"DB_NAME" env:"DB_NAME"`
}

type RedisConfig struct {
	RedisAddress string `yaml:"REDIS_ADDRESS" env:"REDIS_ADDRESS"`
}

func (dc *DbConfig) GetConnectionString() string {
	var connectionString string
	//if strings.Contains(dc.Host, "atlas") {
	//	connectionString = fmt.Sprintf("mongodb+srv://%s", dc.Host)
	//	if dc.Username != "" {
	//		connectionString = fmt.Sprintf("mongodb+srv://%s:%s@%s", dc.Username, dc.Password, dc.Host)
	//	}
	//} else {
	//	connectionString = fmt.Sprintf("mongodb://%s:%s", dc.Host, dc.Port)
	//	if dc.Username != "" {
	//		connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s", dc.Username, dc.Password, dc.Host, dc.Port)
	//	}
	//}
	connectionString = "mongodb://localhost:27017"

	return connectionString
}
