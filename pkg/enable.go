package pkg

import (
	"github.com/NubeDev/bom-api/bom"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	log "github.com/sirupsen/logrus"
)

func (m *Module) Enable() error {
	log.Errorf("enabling plugin %s...", name)
	m.bom = bom.New(&bom.Client{})
	networkUUID, deviceUUID, pointUUID, err := m.getPointByName(networkName, deviceName, pointName)
	if err != nil {
		log.Infof("adding demo point coz: %s", err.Error())
		point, err := m.addDemoPoint(networkName, deviceName, pointName, networkUUID, deviceUUID)
		if err != nil {
			log.Errorf("couldn't add demo points: %s", err.Error())
			return err
		}
		m.demoPointUUID = point.UUID
	} else {
		m.demoPointUUID = *pointUUID
		log.Infof("we already have point with uuid: %s", *pointUUID)
	}
	m.enable = true
	go m.weatherPoll()
	return nil
}

func (m *Module) Disable() error {
	log.Infof("plugin is disabled...%s", name)
	m.enable = false
	return nil
}

// addDemoPoint add new point if not existing
func (m *Module) addDemoPoint(networkName, deviceName, pointName string, networkUUID, deviceUUID *string) (*model.Point, error) {
	if networkUUID == nil {
		network, err := m.grpcMarshaller.CreateNetwork(&model.Network{
			Name:       networkName,
			PluginPath: "system",
		})
		if err != nil {
			return nil, err
		}
		networkUUID = &network.UUID
	}

	if deviceUUID == nil {
		device, err := m.grpcMarshaller.CreateDevice(&model.Device{
			Name:        deviceName,
			NetworkUUID: *networkUUID,
		})
		if err != nil {
			return nil, err
		}
		deviceUUID = &device.UUID
	}

	return m.grpcMarshaller.CreatePoint(&model.Point{
		Name:       pointName,
		DeviceUUID: *deviceUUID,
	})
}
