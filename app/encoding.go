package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	"simapp/app/params"
)

func makeEncodingConfig() params.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interacesRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interacesRegistry)
	txConfig := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	return params.EncodingConfig{
		InterfaceRegistry: interacesRegistry,
		Marshaler:         marshaler,
		TxConfig:          txConfig,
		Amino:             amino,
	}
}

// MakeEncodingConfig creates an EncodingConfig.
func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := makeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
