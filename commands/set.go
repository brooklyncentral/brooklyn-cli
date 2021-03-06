package commands

import (
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_config"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/error_handler"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/codegangsta/cli"
)

type SetConfig struct {
	network *net.Network
}

func NewSetConfig(network *net.Network) (cmd *SetConfig) {
	cmd = new(SetConfig)
	cmd.network = network
	return
}

func (cmd *SetConfig) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "set",
		Description: "Set config for an entity",
		Usage:       "BROOKLYN_NAME CONFIG-SCOPE set VALUE",
		Flags:       []cli.Flag{},
	}
}

func (cmd *SetConfig) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	response, err := entity_config.SetConfig(cmd.network, scope.Application, scope.Entity, scope.Config, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(response)
}
