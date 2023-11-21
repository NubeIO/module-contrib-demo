package pkg

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/args"
	"github.com/NubeIO/rubix-os/module/common"
	log "github.com/sirupsen/logrus"
)

func (m *Module) getPointByName(networkName, deviceName, pointName string) (*string, *string, *string, error) {
	network, err := m.grpcMarshaller.GetNetworkByName(networkName, args.Args{WithDevices: true, WithPoints: true})
	if err != nil {
		log.Error(err)
		return nil, nil, nil, err
	}
	var deviceUUID *string
	for _, device := range network.Devices {
		if device.Name == deviceName {
			deviceUUID = &device.UUID
			for _, point := range device.Points {
				if point.Name == pointName {
					return &network.UUID, &device.UUID, &point.UUID, nil
				}
			}
		}
	}
	if deviceUUID != nil {
		return &network.UUID, deviceUUID, nil, errors.New("point doesn't exist")
	} else {
		return &network.UUID, nil, nil, errors.New("device doesn't exist")
	}
}

func (m *Module) pointWriteAt16(uuid string, value *float64) (*common.PointWriteResponse, error) {
	return m.grpcMarshaller.PointWrite(uuid, writeBody(model.Priority{P16: value}))
}

func (m *Module) pointWrite(uuid string, pointWriter *model.PointWriter) (*common.PointWriteResponse, error) {
	return m.grpcMarshaller.PointWrite(uuid, pointWriter)
}

func (m *Module) getAllPoints(pluginName string) ([]*model.Point, error) {
	get, err := m.grpcMarshaller.GetNetworksByPluginName(pluginName, args.Args{WithDevices: true, WithPoints: true})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var points []*model.Point
	for _, net := range get {
		for _, dev := range net.Devices {
			for _, pnt := range dev.Points {
				points = append(points, pnt)
			}
		}
	}
	return points, err
}

func (m *Module) addPoint(body *model.Point) (*model.Point, error) {
	return m.grpcMarshaller.CreatePoint(body)
}
