package events

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: think about reorganization of blockchain. Check at least several confirmation to be sure that event is valid
const bridgeABI = `[
	{
      "type": "event",
      "name": "Locked",
      "inputs": [
        {
          "name": "token",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "sender",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "cosmosRecipient",
          "type": "string",
          "indexed": false,
          "internalType": "string"
        },
        {
          "name": "amount",
          "type": "uint256",
          "indexed": false,
          "internalType": "uint256"
        },
        {
          "name": "nonce",
          "type": "uint256",
          "indexed": false,
          "internalType": "uint256"
        }
      ],
      "anonymous": false
    }
]`

func Listener() {
	fmt.Println("Start listening...")

	rpcURL := os.Getenv("RPC_WS_URL")
	bridgeAddress := common.HexToAddress(os.Getenv("BRIDGE_ADDRESS"))

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(bridgeABI))
	if err != nil {
		log.Fatal(err)
	}

	event := parsedABI.Events["Locked"]
	query := ethereum.FilterQuery{
		Addresses: []common.Address{bridgeAddress},
		Topics:    [][]common.Hash{{event.ID}},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Println("subscription error:", err)

			case vLog := <-logs:
				var data struct {
					CosmosRecipient string
					Amount          *big.Int
					Nonce           *big.Int
				}

				err := parsedABI.UnpackIntoInterface(&data, "Locked", vLog.Data)
				if err != nil {
					log.Println("decode error:", err)
					continue
				}

				token := common.HexToAddress(vLog.Topics[1].Hex())
				sender := common.HexToAddress(vLog.Topics[2].Hex())

				log.Println("==== Locked Event ====")
				log.Println("Token:", token.Hex())
				log.Println("Sender:", sender.Hex())
				log.Println("CosmosRecipient:", data.CosmosRecipient)
				log.Println("Amount:", data.Amount)
				log.Println("Nonce:", data.Nonce)
			}
		}
	}()
}
