package pkg

import (
	"github.com/NubeDev/bom-api/bom"
	"github.com/NubeIO/rubix-os/module/shared"
	"github.com/patrickmn/go-cache"
	"time"
)

type Module struct {
	dbHelper        shared.DBHelper
	moduleName      string
	grpcMarshaller  shared.Marshaller
	config          *Config
	store           *cache.Cache
	ErrorOnDB       bool
	moduleDirectory string
	bom             *bom.Client
}

func (inst *Module) Init(dbHelper shared.DBHelper, moduleName string) error {
	grpcMarshaller := shared.GRPCMarshaller{DbHelper: dbHelper}
	inst.dbHelper = dbHelper
	inst.moduleName = moduleName
	inst.grpcMarshaller = &grpcMarshaller
	inst.store = cache.New(5*time.Minute, 10*time.Minute)
	dir, err := inst.dbHelper.CreateModuleDataDir(moduleName)
	if err != nil {
		return err
	}
	inst.moduleDirectory = dir
	return nil
}

func (inst *Module) GetInfo() (*shared.Info, error) {
	return &shared.Info{
		Name:       name,
		Author:     "Nube iO",
		Website:    "https://nube-io.com",
		License:    "N/A",
		HasNetwork: false,
	}, nil
}
