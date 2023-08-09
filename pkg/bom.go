package pkg

import (
	"encoding/json"
	"github.com/NubeDev/bom-api/bom"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	log "github.com/sirupsen/logrus"
	"time"
)

func (inst *Module) weatherLoop() ([]byte, *bom.Search, string, error) {
	config := inst.getConfig()
	var currentTemp float64
	var state string
	var town string
	var loopTime time.Duration
	if config != nil {
		loopTime = config.LoopTime
		state = config.State
		town = config.Town
	}
	log.Infof("loop weather for: %s %s every: %d minutes", state, town, loopTime)
	for {
		_, weather, err := inst.getWeather(town, state)
		if err != nil {
			log.Errorf("loop weather: getWeather() err: %s", err.Error())
			return nil, nil, "", err
		}
		if weather != nil {
			currentTemp = weather.Data.Temp
		}
		log.Infof("loop weather for: %s %s currentTemp: %f", state, town, currentTemp)
		writePoint, err := inst.pointWriteAt16(inst.demoPointUUID, nils.NewFloat64(currentTemp))
		if err != nil || writePoint == nil {
			if err != nil {
				log.Errorf("loop weather: pointWriteAt16() err: %s", err.Error())
			}
			log.Errorf("loop weather: pointWriteAt16() failed to write to point on uuid: %s", inst.demoPointUUID)
			return nil, nil, "", err
		}
		log.Infof("loop weather for: %s %s updated point value ok: %f", state, town, currentTemp)
		time.Sleep(1 * time.Minute)
	}
}

func (inst *Module) getWeather(town, state string) ([]byte, *bom.Observations, error) {
	if town == "" {
		town = "Sydney"
	}
	if state == "" {
		state = "NSW"
	}
	log.Infof("GET WEATHER FOR %s %s", town, state)
	get, err := inst.bom.ObservationByTown(town, state)
	if err != nil {
		return nil, nil, err
	}
	data, err := json.Marshal(get)
	if err != nil {
		return nil, nil, err
	}
	return data, get, err
}
