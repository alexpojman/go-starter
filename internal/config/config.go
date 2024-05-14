package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	viper.Viper
}

func NewConfig() *config {
	v := viper.New()

	v.AutomaticEnv()

	return &config{Viper: *v}
}

func (c *config) LoadConfigFile(path, configType, configFile string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", path, configFile)); err == nil {
		c.AddConfigPath(path)
		c.SetConfigType(configType)
		c.SetConfigName(configFile)
		
		err := c.ReadInConfig()

		if err != nil {
			return err
		}
	}

	return nil
}