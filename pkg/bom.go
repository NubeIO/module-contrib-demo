package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/bom-api/bom"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	log "github.com/sirupsen/logrus"
	"time"
)

func (m *Module) weatherPoll() {
	config := m.getConfig()
	var currentTemp float64
	var town string
	var loopTime time.Duration
	if config != nil {
		loopTime = config.LoopTime
		town = config.Town
	}
	log.Infof("polling weather for town=%s in every %d minutes", town, loopTime)
	for {
		if !m.enable {
			return
		}
		_, weather, err := m.getWeather(town)
		if err != nil {
			log.Errorf("polling weather getWeather() err: %s", err.Error())
			time.Sleep(1 * time.Minute)
			continue
		}
		if weather != nil {
			currentTemp = weather.Data.Temp
		}
		log.Infof("polling weather for town=%s, current temp: %f", town, currentTemp)
		_, err = m.pointWriteAt16(m.demoPointUUID, nils.NewFloat64(currentTemp))
		if err != nil {
			log.Errorf("polling weather pointWriteAt16() err: %s", err.Error())
			time.Sleep(1 * time.Minute)
			continue
		}
		log.Infof("polling weather for town=%s, updated point value: %f", town, currentTemp)
		time.Sleep(1 * time.Minute)
	}
}

func (m *Module) getWeather(town string) ([]byte, *bom.Observations, error) {
	if town == "" {
		town = "Sydney"
	}
	log.Infof(fmt.Sprintf("GET weather for town=%s", town))
	get, err := m.bom.ObservationByTown(town)
	if err != nil {
		return nil, nil, err
	}
	data, err := json.Marshal(get)
	if err != nil {
		return nil, nil, err
	}
	return data, get, err
}
