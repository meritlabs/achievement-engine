package utils

import "github.com/btcsuite/btcd/chaincfg"

var mainnetParams = chaincfg.Params{
	Name:             "mainnet",
	PubKeyHashAddrID: 50,
	ScriptHashAddrID: 63,
}

var testnetParams = chaincfg.Params{
	Name:             "testnet",
	PubKeyHashAddrID: 110,
	ScriptHashAddrID: 125,
}

// GetNetParams return network parameters for given network
func GetNetParams(network string) chaincfg.Params {
	if network == "mainnet" {
		return mainnetParams
	}

	return testnetParams
}
