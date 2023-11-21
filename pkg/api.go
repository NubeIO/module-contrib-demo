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

type health struct {
	FormattedDateTime string    `json:"formatted_date_time"`
	TimeDate          time.Time `json:"time_date"`
}

func (m *Module) Get(path string) ([]byte, error) {
	log.Infof("HTTP GET path: %s", path)

	if path == ping { // http://0.0.0.0:1660/api/modules/module-contrib-demo/ping
		return json.Marshal(health{
			FormattedDateTime: time.Now().Format(time.Stamp),
			TimeDate:          time.Now().UTC(),
		})
	}

	if strings.Contains(path, getWeather) { // test endpoint for getting the weather http://0.0.0.0:1660/api/modules/module-contrib-demo/weather/Sydney
		parts := urlSplit(path)
		if len(parts) >= 2 {
			weather, _, err := m.getWeather(parts[1])
			return weather, err
		}
	}
	return nil, errors.New(errNotFound)
}

func (m *Module) Post(path string, body []byte) ([]byte, error) {
	return nil, errors.New(errNotFound)
}

func (m *Module) Put(path, uuid string, body []byte) ([]byte, error) {
	return nil, errors.New(errNotFound)
}

func (m *Module) Patch(path, uuid string, body []byte) ([]byte, error) {
	return nil, errors.New(errNotFound)
}

func (m *Module) Delete(path, uuid string) ([]byte, error) {
	return nil, errors.New(errNotFound)
}
