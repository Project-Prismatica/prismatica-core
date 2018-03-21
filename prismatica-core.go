package main

import (
	"flag"
	"time"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"

	"github.com/Project-Prismatica/prismatica-core/go/prismatica_core/configuration"
)

type programArguments struct {
	AmbassadorConfigurationSource string `env:"CORE_AMBASSADOR_SOURCE_CONFIG_DIR"`
	AmbassadorConfigurationDir string `env:"CORE_AMBASSADOR_CONFIG_DIR"`
	LogDebug bool `env:"CORE_DEBUG" envDefault:"false"`
}

var (
	runtimeArguments = populateProgramArguments()
)

func populateProgramArguments() (args programArguments) {

	env.Parse(&args)

	flag.StringVar(&args.AmbassadorConfigurationDir,
		"ambassador-config-dir",
		args.AmbassadorConfigurationDir,
		"root directory of Ambassador's configuration directory " +
		"[CORE_AMBASSADOR_CONFIG_DIR]",
		)

	flag.StringVar(&args.AmbassadorConfigurationSource,
		"ambassador-config-source-dir",
		args.AmbassadorConfigurationSource,
		"root directory of the sources for Ambassador's configuration " +
		"directory[CORE_AMBASSADOR_SOURCE_CONFIG_DIR]",
	)

	flag.BoolVar(&args.LogDebug,"debug", args.LogDebug,
		"use debug log level [CORE_DEBUG]")

	flag.Parse()

	return
}


func init() {
	if runtimeArguments.LogDebug {
		log.SetLevel(log.DebugLevel)
	}else {
		log.SetLevel(log.InfoLevel)
	}

	if len(runtimeArguments.AmbassadorConfigurationSource) == 0 {
		log.Fatal("must specify ambassador configuration source")
	}
	if len(runtimeArguments.AmbassadorConfigurationDir) == 0 {
		log.Fatal("must specify ambassador configuration directory")
	}
}

func main() {
	log.WithFields(log.Fields{"runtime_arguments": runtimeArguments}).
		Info("prismatica-core starting")

	configuration.HandleAmbassadorConfiguration(
		runtimeArguments.AmbassadorConfigurationSource,
		runtimeArguments.AmbassadorConfigurationDir)

	for true {
		time.Sleep(3600)
	}

	log.Info("prismatica-core shutting down")
}
