package option

import "github.com/urfave/cli"

var Flags = []cli.Flag{
	cli.BoolFlag{
		Name:     "cr",
		Usage:    "is Custom resource",
		Required: false,
	},
	cli.BoolFlag{
		Name:     "namespaced",
		Usage:    "is Custom resource namespaced",
		Required: false,
	},
	cli.StringFlag{
		Name:     "apis",
		Usage:    "Custom resource api def package",
		Required: false,
	},
	cli.StringFlag{
		Name:     "client, c",
		Usage:    "Custom resource client package name",
		Required: false,
	},
	cli.StringFlag{
		Name:     "scheme, s",
		Usage:    "Custom resource scheme package name",
		Required: false,
	},
}
