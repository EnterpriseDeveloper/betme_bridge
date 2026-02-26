package mint

import (
	proto "bridge_betme/proto"
	"context"
	"log"
	"os"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/client/tx"
)

func SendToCosmos(claim Claim, signature []byte) {
	privKey, fromAddr := GetCosmosRelayerAddress()

	msg := &proto.MsgMintFromEvm{
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

	conn, err := grpc.NewClient(
		os.Getenv("COSMOS_GRPC_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	authClient := authtypes.NewQueryClient(conn)

	accRes, err := authClient.Account(
		context.Background(),
		&authtypes.QueryAccountRequest{
			Address: fromAddr.String(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	var account authtypes.BaseAccount
	err = account.Unmarshal(accRes.Account.Value)
	if err != nil {
		log.Fatal(err)
	}

	accountNumber := account.AccountNumber
	sequence := account.Sequence

	// --- encoding ---
	txConfig := MakeEncodingConfig()
	txBuilder := txConfig.TxConfig.NewTxBuilder()

	err = txBuilder.SetMsgs(msg)
	if err != nil {
		log.Fatal(err)
	}

	txBuilder.SetGasLimit(500000) // TODO: change

	signMode := signingtypes.SignMode_SIGN_MODE_DIRECT

	sigV2 := signingtypes.SignatureV2{
		PubKey: privKey.PubKey(),
		Data: &signingtypes.SingleSignatureData{
			SignMode:  signMode,
			Signature: nil,
		},
		Sequence: sequence,
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		log.Fatal(err)
	}

	signerData := authsigning.SignerData{
		ChainID:       os.Getenv("COSMOS_CHAIN_ID"),
		AccountNumber: accountNumber,
		Sequence:      sequence,
	}

	sigV2, err = tx.SignWithPrivKey(
		context.Background(),
		signMode,
		signerData,
		txBuilder,
		privKey,
		txConfig.TxConfig,
		sequence,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		log.Fatal(err)
	}

	// --- encode ---
	txBytes, err := txConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		log.Fatal(err)
	}

	// --- broadcast ---
	txClient := txtypes.NewServiceClient(conn)

	res, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Cosmos Tx Hash:", res.TxResponse.TxHash)

}
