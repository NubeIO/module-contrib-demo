package main

import (
	"github.com/NubeIO/lib-module-go/nmodule"
	"github.com/NubeIO/module-contrib-demo/pkg"
	"github.com/hashicorp/go-plugin"
)

const name = "module-contrib-demo"

func ServePlugin() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: nmodule.HandshakeConfig,
		Plugins:         plugin.PluginSet{name: &nmodule.NubeModule{Impl: &pkg.Module{}}},
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}

func main() {
	ServePlugin()
}
