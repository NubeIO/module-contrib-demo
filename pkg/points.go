package pkg

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bugs"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/args"
	"github.com/NubeIO/rubix-os/module/common"
	log "github.com/sirupsen/logrus"
)

func (inst *Module) getPointByName(networkName, deviceName, pointName string) (*model.Point, error) {
	get, err := inst.grpcMarshaller.GetNetworkByName(networkName, args.Args{WithDevices: true, WithPoints: true})
	if err != nil {
		log.Error(bugs.DebugPrint(name, inst.getAllPoints, err))
		return nil, err
	}
	var point *model.Point
	for _, dev := range get.Devices {
		if dev.Name == deviceName {
			for _, pnt := range dev.Points {
				if pnt.Name == pointName {
					point = pnt
				}
			}
		}
	}
	return point, err
}

func (inst *Module) pointWriteAt16(uuid string, value *float64) (*common.PointWriteResponse, error) {
	return inst.grpcMarshaller.PointWrite(uuid, writeBody(model.Priority{P16: value}))
}

func (inst *Module) pointWrite(uuid string, pointWriter *model.PointWriter) (*common.PointWriteResponse, error) {
	return inst.grpcMarshaller.PointWrite(uuid, pointWriter)
}

func (inst *Module) getAllPoints(pluginName string) ([]*model.Point, error) {
	get, err := inst.grpcMarshaller.GetNetworksByPluginName(pluginName, args.Args{WithDevices: true, WithPoints: true})
	if err != nil {
		log.Error(bugs.DebugPrint(name, inst.getAllPoints, err))
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
