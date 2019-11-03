package config

import (
	"strings"

	"go.uber.org/zap"

	"github.com/alexsniffin/go-places/internal/api/models"
	"github.com/spf13/viper"
)

//Config todo
type Config struct {
	Cfg *models.Configuration
}

const (
	envPrefix       = "GO_PLACES"
	localConfigPath = "$GOPATH/src/github.com/alexsniffin/go-places/configs/"
)

//NewConfig todo
func NewConfig(log *zap.SugaredLogger, filename string) *Config {
	v, err := readConfig(filename, map[string]interface{}{})
	if err != nil {
		log.Panic("Failed to initialize config for file " + filename)
	}

	var cfg = &models.Configuration{}
	err = v.Unmarshal(cfg)
	if err != nil {
		log.Panic("Failed to decode config values")
	}

	return &Config{
		Cfg: cfg,
	}
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(localConfigPath) // check the abs path for running locally
	v.AddConfigPath(".")             // check alongside the binary for deployments
	v.SetConfigType("yaml")
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
