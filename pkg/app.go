package pkg

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	log "github.com/sirupsen/logrus"
	"time"
)

const networkName = "demo-network"
const deviceName = "demo-device"
const pointName = "demo-point"

func (m *Module) networkUpdateSuccess(uuid string) error {
	var network model.Network
	network.InFault = false
	network.MessageLevel = dto.MessageLevel.Info
	network.MessageCode = dto.CommonFaultCode.Ok
	network.Message = dto.CommonFaultMessage.NetworkMessage
	network.LastOk = time.Now().UTC()
	err := m.grpcMarshaller.UpdateNetworkErrors(uuid, &network)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (m *Module) networkUpdateErr(uuid, port string, e error) error {
	var network model.Network
	network.InFault = true
	network.MessageLevel = dto.MessageLevel.Fail
	network.MessageCode = dto.CommonFaultCode.NetworkError
	network.Message = fmt.Sprintf(" port: %s message: %s", port, e.Error())
	network.LastFail = time.Now().UTC()
	err := m.grpcMarshaller.UpdateNetworkErrors(uuid, &network)
	if err != nil {
		log.Error(err)
	}
	return err
}
