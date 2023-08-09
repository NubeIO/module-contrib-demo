package pkg

import (
	"github.com/NubeDev/bom-api/bom"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	log "github.com/sirupsen/logrus"
)

func (inst *Module) Enable() error {
	log.Error("plugin is enabling...%s", name)
	// init BOM client
	inst.bom = bom.New(&bom.Client{})

	// add new point for demo
	point, err := inst.getPointByName(nameNetwork, nameDevice, namePoint)
	if err != nil {
		log.Errorf("adding network: no existing point: %s", err.Error())
		point, err = inst.addDemoPoint("", nameNetwork, nameDevice, namePoint)
		if err != nil {
			log.Errorf("adding network: error, now try and get existing point: %s", err.Error())
			return err
		}
		inst.demoPointUUID = point.UUID
		log.Infof("point was existing, uuid: %s", point.UUID)
	} else {
		inst.demoPointUUID = point.UUID
		log.Infof("adding new point ok uuid: %s", point.UUID)
	}
	go inst.weatherLoop()
	return nil
}

func (inst *Module) Disable() error {
	log.Infof("plugin is disabled...%s", name)
	return nil
}

// addDemoPoint add new point if not existing, if pluginName is "" it will use the system plugin
func (inst *Module) addDemoPoint(pluginName, networkName, deviceName, pointName string) (*model.Point, error) {
	if pluginName == "" {
		pluginName = "system"
	}
	var err error
	network, err := inst.grpcMarshaller.CreateNetwork(&model.Network{
		Name:       networkName,
		PluginPath: pluginName,
	})
	if err != nil {
		return nil, err
	}
	device, err := inst.grpcMarshaller.CreateDevice(&model.Device{
		Name:        deviceName,
		NetworkUUID: network.UUID,
	})
	if err != nil {
		return nil, err
	}
	point, err := inst.grpcMarshaller.CreatePoint(&model.Point{
		Name:       pointName,
		DeviceUUID: device.UUID,
	})

	return point, err
}
