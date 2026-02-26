package mint

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	fmt.Println("Start listening EVM...")

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

				bridgeAddr := common.HexToAddress(os.Getenv("BRIDGE_ADDRESS"))
				evmChainID := os.Getenv("EVM_CHAIN_ID")
				evmChainIDnum, err := strconv.ParseUint(evmChainID, 10, 64)
				if err != nil {
					log.Println("parse EVM_CHAIN_ID:", err)
				}

				claim := Claim{
					EvmChainId:     evmChainIDnum, // Polygon Amoy
					EvmBridge:      bridgeAddr,
					EvmToken:       token,
					EvmSender:      sender,
					CosmosReceiver: data.CosmosRecipient,
					Amount:         data.Amount,
					Nonce:          data.Nonce.Uint64(),
					TxHash:         vLog.TxHash,
				}

				hash := HashClaim(claim)

				log.Println("Claim hash:", common.BytesToHash(hash).Hex())

				// ---- SIGN ----

				privateKeyHex := os.Getenv("RELAYER_EVM_PRIVATE_KEY")
				pk, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
				if err != nil {
					log.Fatal(err)
				}

				signature, err := SignClaim(hash, pk)
				if err != nil {
					log.Fatal(err)
				}

				log.Println("Signature:", hex.EncodeToString(signature))

				SendToCosmos(claim, signature)

			}
		}
	}()
}
