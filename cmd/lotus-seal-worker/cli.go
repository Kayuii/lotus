package main

import (
	"github.com/filecoin-project/lotus/cli/util"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var setCmd = &cli.Command{
	Name:  "set",
	Usage: "Manage worker settings",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "enabled",
			Usage: "enable/disable new task processing",
			Value: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := cliutil.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := cliutil.ReqContext(cctx)

		if err := api.SetEnabled(ctx, cctx.Bool("enabled")); err != nil {
			return xerrors.Errorf("SetEnabled: %w", err)
		}

		return nil
	},
}

var waitQuietCmd = &cli.Command{
	Name:  "wait-quiet",
	Usage: "Block until all running tasks exit",
	Action: func(cctx *cli.Context) error {
		api, closer, err := cliutil.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := cliutil.ReqContext(cctx)

		return api.WaitQuiet(ctx)
	},
}
