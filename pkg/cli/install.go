package cli

import (
	"fmt"

	"github.com/aquaproj/aqua/pkg/controller"
	"github.com/urfave/cli/v2"
)

func (runner *Runner) setCLIArg(c *cli.Context, param *controller.Param) error { //nolint:unparam
	if logLevel := c.String("log-level"); logLevel != "" {
		param.LogLevel = logLevel
	}
	param.ConfigFilePath = c.String("config")
	param.OnlyLink = c.Bool("only-link")
	param.IsTest = c.Bool("test")
	param.All = c.Bool("all")
	param.File = c.String("f")
	param.AQUAVersion = runner.LDFlags.Version
	return nil
}

func (runner *Runner) installAction(c *cli.Context) error {
	param := &controller.Param{}
	if err := runner.setCLIArg(c, param); err != nil {
		return fmt.Errorf("parse the command line arguments: %w", err)
	}

	ctrl, err := controller.New(c.Context, param)
	if err != nil {
		return fmt.Errorf("initialize a controller: %w", err)
	}

	return ctrl.Install(c.Context, param) //nolint:wrapcheck
}
