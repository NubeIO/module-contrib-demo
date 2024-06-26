package pkg

import (
	"github.com/NubeIO/bom-api/bom"
	"github.com/NubeIO/lib-module-go/nmodule"
	"github.com/NubeIO/lib-utils-go/nstring"
	"github.com/patrickmn/go-cache"
	"time"
)

type Module struct {
	dbHelper        nmodule.DBHelper
	moduleName      string
	grpcMarshaller  nmodule.Marshaller
	config          *Config
	store           *cache.Cache
	ErrorOnDB       bool
	moduleDirectory string
	bom             *bom.Client
	enable          bool
	demoPointUUID   string
}

func (m *Module) Init(dbHelper nmodule.DBHelper, moduleName string) error {
	InitRouter()
	grpcMarshaller := nmodule.GRPCMarshaller{DbHelper: dbHelper}
	m.dbHelper = dbHelper
	m.moduleName = moduleName
	m.grpcMarshaller = &grpcMarshaller
	m.store = cache.New(5*time.Minute, 10*time.Minute)
	dir, err := m.grpcMarshaller.CreateModuleDir(moduleName)
	if err != nil {
		return err
	}
	m.moduleDirectory = nstring.DerefString(dir)
	return nil
}

func (m *Module) GetInfo() (*nmodule.Info, error) {
	return &nmodule.Info{
		Name:       m.moduleName,
		Author:     "Nube iO",
		Website:    "https://nube-io.com",
		License:    "N/A",
		HasNetwork: false,
	}, nil
}
