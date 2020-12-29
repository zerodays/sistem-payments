package config

import (
	"github.com/zerodays/gonfig"
	"github.com/zerodays/gonfig/providers/environment"
	"github.com/zerodays/gonfig/providers/ini"
	"log"
)

var cfg *gonfig.Config

// Load loads config
func Load() {
	// Create ini provider with default config.
	data, err := Asset("config.ini")
	if err != nil {
		log.Fatal(err)
	}

	iniProvider, err := ini.New(data)
	if err != nil {
		log.Fatal(err)
	}
	iniProvider.SectionMode = ini.ModeLowercase
	iniProvider.KeyMode = ini.ModeUppercase

	// Create gonfig instance.
	cfg = gonfig.New(iniProvider, environment.Provider{})
	cfg.AppName = "sistem_projects"
}
