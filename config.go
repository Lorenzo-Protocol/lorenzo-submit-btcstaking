package main

import (
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"

	lrzcfg "github.com/Lorenzo-Protocol/lorenzo-sdk/config"
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string               `mapstructure:"logLevel"`
	Lorenzo  lrzcfg.LorenzoConfig `mapstructure:"lorenzo"`
	BTC      BTCConfig            `mapstructure:"btc"`
}

type BTCConfig struct {
	PreHandledTxid       string `mapstructure:"preHandledTxid"`
	ConfirmationDepth    int    `mapstructure:"confirmationDepth"`
	NetParams            string `mapstructure:"netParams"`
	TargetDepositAddress string `mapstructure:"targetDepositAddress"`
}

func (cfg *BTCConfig) Validate() error {
	if cfg.ConfirmationDepth <= 0 {
		return errors.New("confirmationDepth must be positive")
	}
	if cfg.NetParams != "mainnet" && cfg.NetParams != "testnet" {
		return errors.New("netParams must be either mainnet or testnet")
	}
	if cfg.TargetDepositAddress == "" {
		return errors.New("targetDepositAddress must be set")
	}

	return nil
}

func (cfg *Config) Validate() error {
	if err := cfg.Lorenzo.Validate(); err != nil {
		return err
	}

	if err := cfg.BTC.Validate(); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) IsTestNet() bool {
	return cfg.BTC.NetParams == "testnet"
}

func (cfg *Config) CreateLogger() (*zap.Logger, error) {
	return newRootLogger("auto", cfg.LogLevel == "debug")
}

// NewConfig returns a fully parsed Config object from a given file directory
func NewConfig(configFile string) (Config, error) {
	if _, err := os.Stat(configFile); err == nil { // the given file exists, parse it
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			return Config{}, err
		}
		var cfg Config
		if err := viper.Unmarshal(&cfg); err != nil {
			return Config{}, err
		}
		if err := cfg.Validate(); err != nil {
			return Config{}, err
		}
		return cfg, err
	} else if errors.Is(err, os.ErrNotExist) { // the given config file does not exist, return error
		return Config{}, fmt.Errorf("no config file found at %s", configFile)
	} else { // other errors
		return Config{}, err
	}
}