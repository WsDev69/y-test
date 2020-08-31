package env

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"y-test/internal/module/config"

	"os"
	"sync"
)

type Config struct {
	v *config.Values
}

var instantiated *Config
var once sync.Once

func NewEnvConfig() config.Config {
	once.Do(func() {
		var values config.Values
		var c = &Config{}
		c.ReadConfig(&values)
		c.v = &values
		instantiated = c
	})
	return instantiated
}

func (c *Config) ReadConfig(i interface{}) {
	err := envconfig.Process("", i)
	if err != nil {
		logrus.WithError(err).Error("envconfig couldn't read conf data")
		os.Exit(2)
	}
}

func (c *Config) GetValues() *config.Values {
	return c.v
}
