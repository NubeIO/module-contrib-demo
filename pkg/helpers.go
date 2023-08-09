package pkg

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/utils/float"
	"reflect"
)

func writeBody(priority model.Priority) *model.PointWriter {
	pri := convertPriority(priority)
	body := &model.PointWriter{
		Priority:     &pri,
		PresentValue: nil,
		ForceWrite:   false,
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
