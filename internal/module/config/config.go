package config

type Config interface {
	ReadConfig(c interface{})
	GetValues() *Values
}

type Values struct {
	Sqlite struct {
		Path string `envconfig:"SQLITE_PATH"`
	}
	Server struct {
		Port int32 `envconfig:"SERVER_PORT"`
	}
	JWT struct {
		Secret string `envconfig:"JWT_SECRET"`
	}
	Storage struct {
		Endpoint        string `enconfig:"STORAGE_ENDPOINT"`
		AccessKeyId     string `enconfig:"STORAGE_ACCESS_KEY_ID"`
		SecretAccessKey string `enconfig:"STORAGE_SECRET_ACCESS_KEY"`
		UseSSL          bool   `enconfig:"STORAGE_USESSL"`
	}
}
