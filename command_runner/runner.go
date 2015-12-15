package command_runner

import (
    "fmt"
    "github.com/codegangsta/cli"
    "github.com/brooklyncentral/brooklyn-cli/command_factory"
    "github.com/brooklyncentral/brooklyn-cli/scope"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

type Runner interface {
    RunCmdByName(cmdName string, c *cli.Context) (err error)
    RunSubCmdByName(cmdName string, subCommand string, c *cli.Context) (err error)
}

type ConcreteRunner struct {
    cmdFactory command_factory.Factory
    scope      scope.Scope
}

func NewRunner(scope scope.Scope, cmdFactory command_factory.Factory) (runner ConcreteRunner) {
    runner.cmdFactory = cmdFactory
    runner.scope = scope
    return
}

func (runner ConcreteRunner) RunCmdByName(cmdName string, c *cli.Context) error {
    cmd, err := runner.cmdFactory.GetByCmdName(cmdName)
    if nil != err {
        fmt.Println(err)
        return err
    }

    if err := cmd.Verify(); err != nil {
        error_handler.ErrorExit(err)
    }
    cmd.Run(runner.scope, c)
    return nil
}

func (runner ConcreteRunner) RunSubCmdByName(cmdName string, subCommand string, c *cli.Context) error {
    cmd, err := runner.cmdFactory.GetBySubCmdName(cmdName, subCommand)
    if nil != err {
        fmt.Println(err)
        return err
    }

    cmd.Run(runner.scope, c)
    return nil
}
