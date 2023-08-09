package pkg

import (
	"encoding/json"
	"github.com/NubeDev/bom-api/bom"
)

func (inst *Module) getWeather(town, state string) ([]byte, *bom.Search, string, error) {

	if town == "" {
		town = "Sydney"
	}
	if state == "" {
		state = "NSW"
	}

	search, geo, err := inst.bom.SearchByTown(town, state)
	if err != nil {
		return nil, nil, "", err
	}
	data, err := json.Marshal(search)
	if err != nil {
		return nil, nil, "", err
	}
	return data, search, geo, err

}
