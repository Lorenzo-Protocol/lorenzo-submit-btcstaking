package main

import (
	"flag"

	lrzclient "github.com/Lorenzo-Protocol/lorenzo-sdk/client"

	"github.com/Lorenzo-Protocol/lorenzo-btcstaking-submitter/btc"
	"github.com/Lorenzo-Protocol/lorenzo-btcstaking-submitter/config"
	"github.com/Lorenzo-Protocol/lorenzo-btcstaking-submitter/db"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "./.testnet/sample-config.yml", "config file")
	flag.Parse()

	cfg, err := config.NewConfig(configFile)
	if err != nil {
		panic(err)
	}
	database, err := db.NewMysqlDB(cfg.Database)
	if err != nil {
		panic(err)
	}

	lorenzoClient, err := lrzclient.New(&cfg.Lorenzo, nil)
	if err != nil {
		panic(err)
	}

	btcQuery := btc.NewBTCQuery(cfg.BtcApiEndpoint)

	parentLogger, err := cfg.CreateLogger()
	if err != nil {
		panic(err)
	}
	logger := parentLogger.With().Sugar()

	txRelayer, err := NewTxRelayer(database, logger, &cfg.TxRelayer, btcQuery, lorenzoClient)
	if err != nil {
		panic(err)
	}
	if err := txRelayer.Start(); err != nil {
		panic(err)
	}
}
