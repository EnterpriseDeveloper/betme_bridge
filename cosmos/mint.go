package cosmos

import (
	helper "bridge_betme/helper"
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func SendToCosmos(claim helper.Claim, signature []byte) {
	privKey, fromAddr := GetCosmosRelayerAddress()

	msg := &MsgMintFromEvm{
		Creator:        fromAddr.String(),
		EvmChainId:     claim.EvmChainId,
		EvmBridge:      claim.EvmBridge.Hex(),
		EvmToken:       claim.EvmToken.Hex(),
		EvmSender:      claim.EvmSender.Hex(),
		CosmosReceiver: claim.CosmosReceiver,
		Amount:         claim.Amount.String(),
		Nonce:          claim.Nonce,
		TxHash:         claim.TxHash.Hex(),
		Signatures:     [][]byte{signature},
	}

	txConfig := MakeEncodingConfig()
	// --- создаём keyring из mnemonic ---
	kr := BuildKeyringFromPriv(privKey, txConfig.Codec)

	// --- client context ---
	clientCtx := client.Context{}.
		WithChainID(os.Getenv("COSMOS_CHAIN_ID")).
		WithTxConfig(txConfig.TxConfig).
		WithKeyring(kr).
		WithFromName("alice").
		WithFromAddress(fromAddr).
		WithNodeURI(os.Getenv("COSMOS_RPC_URL")).
		WithBroadcastMode("sync")

	// --- tx factory ---
	txFactory := tx.Factory{}.
		WithTxConfig(txConfig.TxConfig).
		WithChainID(os.Getenv("COSMOS_CHAIN_ID")).
		WithGas(500000)

	// --- автоматическая подпись + broadcast ---
	err := tx.GenerateOrBroadcastTxWithFactory(
		clientCtx,
		txFactory,
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Mint tx broadcasted successfully")

}
