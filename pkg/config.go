package pkg

import (
	"github.com/NubeIO/module-contrib-demo/logger"
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Config struct {
	Town     string        `yaml:"town"`
	LoopTime time.Duration `yaml:"loop_time"`
	LogLevel string        `yaml:"log_level"`
}

func (m *Module) DefaultConfig() *Config {
	return &Config{
		Town:     "Sydney",
		LoopTime: 1,
		LogLevel: "INFO", // INFO, DEBUG, ERROR
	}
}

func (m *Module) getConfig() *Config {
	return m.config
}

func (m *Module) GetConfig() interface{} {
	return m.config
}

func (m *Module) ValidateAndSetConfig(config []byte) ([]byte, error) {
	newConfig := m.DefaultConfig()
	_ = yaml.Unmarshal(config, newConfig) // if unable to marshal just take the default one

	logLevel, err := log.ParseLevel(newConfig.LogLevel)
	if err != nil {
		logLevel = log.ErrorLevel
	}
	logger.SetLogger(logLevel)

	newConfig.LogLevel = strings.ToUpper(logLevel.String())

	newConfValid, err := yaml.Marshal(newConfig)
	if err != nil {
		return nil, err
	}
	m.config = newConfig

	log.Info("config is set")
	return newConfValid, nil
}
