package main

import (
	"github.com/urfave/cli/v2"
	"github.com/zerodays/sistem-auth/token"
	"github.com/zerodays/sistem-payments/internal/cmd"
	"github.com/zerodays/sistem-payments/internal/config"
	"github.com/zerodays/sistem-payments/internal/database"
	"github.com/zerodays/sistem-payments/internal/logger"
	"os"
)

func main() {
	// Load configuration.
	config.Load()

	// Initialize logger instance.
	logger.Init()

	// Initialize database.
	database.Init()
	defer database.Close()

	// Loads RSA public signing key for user authentication
	err := token.LoadKey(config.Microservices.UsersUrl() + "/signing_key")
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}

	// Create new CLI app.
	app := cli.NewApp()

	// Basic info.
	app.Name = "Sistem payments microservice"
	app.Authors = []*cli.Author{
		{
			Name:  "Vid Drobnič",
			Email: "vid.drobnic@gmail.com",
		},
	}
	app.Version = "0.0.1"

	// Commands for CLI app.
	app.Commands = []*cli.Command{
		cmd.Serve,
	}

	// Run the app.
	err = app.Run(os.Args)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}
}
