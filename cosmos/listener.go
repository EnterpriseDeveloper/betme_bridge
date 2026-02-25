package cosmos

import (
	"context"
	"fmt"
	"log"
	"os"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cometbft/cometbft/types"
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
				}
			}
		}
	}()
}
