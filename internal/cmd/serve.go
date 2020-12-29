package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/zerodays/sistem-payments/internal/config"
	"github.com/zerodays/sistem-payments/internal/logger"
	"github.com/zerodays/sistem-payments/internal/router"
	"net/http"
)

var Serve = &cli.Command{
	Name:   "serve",
	Usage:  "Start the sever.",
	Action: serve,
}

func serve(_ *cli.Context) error {
	// Load router.
	r := router.NewRouter()

	// Start listening for connections.
	listenAddress := fmt.Sprintf("%s:%d", config.Server.ListenAddress(), config.Server.Port())
	logger.Log.Fatal().Err(http.ListenAndServe(listenAddress, r)).Send()

	return nil
}
