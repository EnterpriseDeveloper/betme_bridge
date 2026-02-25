package cosmos

import (
	helper "bridge_betme/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func SendToCosmos(claim helper.Claim, signature []byte) {

	msg := &MsgMintFromEvm{
		Creator:        cosmosRelayerAddress,
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

	cosmosURL := os.Getenv("RELAYER_EVM_PRIVATE_KEY")
	conn, err := grpc.NewClient(
		cosmosURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}

	txClient := txtypes.NewServiceClient(conn)
}
