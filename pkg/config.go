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
	State    string        `yaml:"state"`
	LoopTime time.Duration `yaml:"loop_time"`
	LogLevel string        `yaml:"log_level"`
}

func (inst *Module) DefaultConfig() *Config {
	return &Config{
		Town:     "Sydney",
		State:    "NSW",
		LoopTime: 1,
		LogLevel: "INFO", // INFO, DEBUG, ERROR
	}
}

func (inst *Module) getConfig() *Config {
	return inst.config
}

func (inst *Module) GetConfig() interface{} {
	return inst.config
}

func (inst *Module) ValidateAndSetConfig(config []byte) ([]byte, error) {
	newConfig := inst.DefaultConfig()
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
	inst.config = newConfig

	log.Info("config is set")
	return newConfValid, nil
}
