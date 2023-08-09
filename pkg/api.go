package pkg

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	jsonSchemaNetwork = "/schema/json/network"
	jsonSchemaDevice  = "/schema/json/device"
	jsonSchemaPoint   = "/schema/json/point"
	getPoints         = "/points"
	getWeather        = "/weather"
)

const errNotFound = "not found"

func urlSplit(path string) []string {
	return strings.Split(path, "/")
}

func urlLen(path string) int {
	return len(strings.Split(path, "/"))
}

func urlIsCorrectModule(path string) bool {
	for _, s := range urlSplit(path) {
		if s == name {
			return true
		}
	}
	return false
}

func (inst *Module) Get(path string) ([]byte, error) {
	log.Error(11111, path)
	if strings.Contains(path, getWeather) { // test endpoint for getting the weather http://0.0.0.0:1660/api/modules/module-contrib-demo/weather/Sydney/NSW
		parts := urlSplit(path)
		if len(parts) == 3 {
			weather, _, _, err := inst.getWeather(parts[1], parts[2])
			if err != nil {
				return weather, err
			}
		}
	}

	return nil, errors.New(path)
}

func (inst *Module) Post(path string, body []byte) ([]byte, error) {

	return nil, errors.New(errNotFound)
}

func (inst *Module) Put(path, uuid string, body []byte) ([]byte, error) {
	return nil, errors.New(errNotFound)
}

func (inst *Module) Patch(path, uuid string, body []byte) ([]byte, error) {

	return nil, errors.New(errNotFound)
}

func (inst *Module) Delete(path, uuid string) ([]byte, error) {

	return nil, errors.New(errNotFound)
}
