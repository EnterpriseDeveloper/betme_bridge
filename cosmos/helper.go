package cosmos

import (
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"

	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
}

func GetCosmosRelayerAddress() (*secp256k1.PrivKey, sdk.AccAddress) {
	mnemonic := os.Getenv("COSMOS_RELAYER_MNEMONIC")

	if !bip39.IsMnemonicValid(mnemonic) {
		log.Fatal("invalid mnemonic")
	}

	// Cosmos derivation path
	hdPath := hd.CreateHDPath(118, 0, 0).String()

	masterPriv, ch := hd.ComputeMastersFromSeed(bip39.NewSeed(mnemonic, ""))

	derivedPriv, err := hd.DerivePrivateKeyForPath(masterPriv, ch, hdPath)
	if err != nil {
		log.Fatal(err)
	}

	privKey := &secp256k1.PrivKey{
		Key: derivedPriv,
	}

	fromAddr := sdk.AccAddress(privKey.PubKey().Address())

	log.Println("Cosmos relayer address:", fromAddr.String())
	return privKey, fromAddr
}

func MakeEncodingConfig() EncodingConfig {
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	std.RegisterInterfaces(interfaceRegistry)

	RegisterInterfaces(interfaceRegistry)

	protoCodec := codec.NewProtoCodec(interfaceRegistry)

	txConfig := authtx.NewTxConfig(
		protoCodec,
		authtx.DefaultSignModes,
	)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             protoCodec,
		TxConfig:          txConfig,
	}
}

func BuildKeyringFromPriv(priv *secp256k1.PrivKey, cdc codec.Codec) keyring.Keyring {

	kr := keyring.NewInMemory(cdc)

	err := kr.ImportPrivKey(
		"alice",
		hex.EncodeToString(priv.Bytes()),
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	return kr
}
