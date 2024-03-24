package pkg

import (
	"github.com/NubeIO/lib-utils-go/float"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"reflect"
)

func writeBody(priority model.Priority) *dto.PointWriter {
	pri := convertPriority(priority)
	body := &dto.PointWriter{
		Priority:   &pri,
		ForceWrite: false,
	}
	return body
}

func convertPriority(priority model.Priority) map[string]*float64 {
	priorityMap := map[string]*float64{}
	priorityValue := reflect.ValueOf(priority)
	typeOfPriority := priorityValue.Type()
	for i := 0; i < priorityValue.NumField(); i++ {
		if priorityValue.Field(i).Type().Kind().String() == "ptr" {
			key := typeOfPriority.Field(i).Tag.Get("json")
			val := priorityValue.Field(i).Interface().(*float64)
			var val64 *float64
			if val == nil {
				val64 = nil
			} else {
				val64 = float.New(*val)
			}
			priorityMap[key] = val64
		}
	}
	return priorityMap
}
