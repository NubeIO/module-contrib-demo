package pkg

import (
	"github.com/NubeDev/bom-api/bom"
	log "github.com/sirupsen/logrus"
)

func (inst *Module) Enable() error {
	log.Error("plugin is enabling...%s", name)

	log.Infof("plugin is enabled...%s", name)

	inst.bom = bom.New(&bom.Client{})

	return nil
}

func (inst *Module) Disable() error {
	log.Infof("plugin is disabled...%s", name)
	return nil
}
