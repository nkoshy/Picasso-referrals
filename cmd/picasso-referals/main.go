package main

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/xlab/suplog"

	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
)

var app = cli.App("picasso-referal", "picasso-referal is a executable for picasso referal service")

var (

	// General config
	envName     *string
	appLogLevel *string
)

func main() {
	readEnv()
	app.Before = prepareApp
	initGlobalOptions(
		&envName,
		&appLogLevel,
	)

	fmt.Println("picasso-referal config", "envname", *envName, "appLogLevel", *appLogLevel)

	app.Command("refer-user", "Start user referals ", orchestratorCmd)

	_ = app.Run(os.Args)
}

func prepareApp() {
	config := cosmtypes.GetConfig()
	chaintypes.SetBech32Prefixes(config)
	chaintypes.SetBip44CoinType(config)
}

func orShutdown(err error) {
	if err != nil {
		log.WithError(err).Fatalln("unable to start picasso bridge")
	}
}
