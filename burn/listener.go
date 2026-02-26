package burn

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"encoding/hex"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func StartCosmosListener() {
	fmt.Println("Start listening COSMOS...")

	client, err := rpchttp.New(os.Getenv("COSMOS_WEBSOCKET_URL"), "/websocket")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Start()
	if err != nil {
		log.Fatal(err)
	}

	query := "tm.event='Tx' AND burn_to_evm.nonce EXISTS"

	ctx := context.Background()

	sub, err := client.Subscribe(ctx, "bridge-relayer", query)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyRelayer1Hex := os.Getenv("RELAYER_EVM_PRIVATE_KEY")
	relayer1, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyRelayer1Hex, "0x"))
	if err != nil {
		log.Fatal(err)
	}

	privateKeyRelayer2Hex := os.Getenv("RELAYER_TWO_EVM_PRIVATE_KEY")
	relayer2, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyRelayer2Hex, "0x"))
	if err != nil {
		log.Fatal(err)
	}

	privateKeyRelayer3Hex := os.Getenv("RELAYER_THREE_EVM_PRIVATE_KEY")
	relayer3, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyRelayer3Hex, "0x"))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range sub {

			txResult := msg.Data.(types.EventDataTx)

			log.Println("New burn event!")

			for _, event := range txResult.Result.Events {

				if event.Type == "burn_to_evm" {

					for _, attr := range event.Attributes {
						log.Printf("%s: %s\n",
							attr.Key,
							attr.Value,
						)
					}

					claim := ParseBurnEvent(event)

					hash := HashBurnClaim(claim)

					signatureRel1, _ := SignClaim(hash, relayer1)
					signatureRel2, _ := SignClaim(hash, relayer2)
					signatureRel3, _ := SignClaim(hash, relayer3)

					log.Println("Hash:", hex.EncodeToString(hash))
					log.Println("Relayer 1:", crypto.PubkeyToAddress(relayer1.PublicKey).Hex())
					log.Println("Relayer 2:", crypto.PubkeyToAddress(relayer2.PublicKey).Hex())
					log.Println("Relayer 3:", crypto.PubkeyToAddress(relayer3.PublicKey).Hex())

					// TODO: need test, not sure that it works
					SendUnlockToEVM(claim, signatureRel1, signatureRel2, signatureRel3)

				}
			}
		}
	}()
}
