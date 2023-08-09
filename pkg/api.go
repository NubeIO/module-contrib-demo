package pkg

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	jsonSchemaNetwork = "/schema/json/network"
	jsonSchemaDevice  = "/schema/json/device"
	jsonSchemaPoint   = "/schema/json/point"
	ping              = "/ping"
	getPoints         = "/points"
	getWeather        = "/weather"
)

const errNotFound = "not found"

func urlSplit(path string) []string {
	var out []string
	parts := strings.Split(path, "/")
	for i, part := range parts {
		log.Infof("part: %s %d count: %d", part, len(part), i)
		if len(part) > 0 {
			out = append(out, part)
		}
	}
	return out
}

func urlIsCorrectModule(path string) bool {
	for _, s := range urlSplit(path) {
		if s == name {
			return true
		}
	}
	return false
}

type helloWorld struct {
	A              string    `json:"a"`
	B              int       `json:"b"`
	C              bool      `json:"c"`
	TimeDateFormat string    `json:"time_date_format"`
	TimeDate       time.Time `json:"time_date"`
}

func (inst *Module) Get(path string) ([]byte, error) {
	log.Infof("HTTP-GET path: %s", path)
	//log.Info("HTTP-GET path:", strings.Contains(path, getWeather))
	if path == ping { //http://0.0.0.0:1660/api/modules/module-contrib-demo/ping
		return json.Marshal(helloWorld{
			A:              "ping",
			B:              0,
			C:              false,
			TimeDateFormat: time.Now().Format(time.Stamp),
			TimeDate:       time.Now().UTC(),
		})
	}

	if strings.Contains(path, getWeather) { // test endpoint for getting the weather http://0.0.0.0:1660/api/modules/module-contrib-demo/weather/Sydney/NSW
		parts := urlSplit(path)
		if len(parts) >= 3 {
			weather, _, err := inst.getWeather(parts[1], parts[2])
			return weather, err
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
