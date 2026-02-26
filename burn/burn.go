package burn

import (
	abi "bridge_betme/abi"
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendUnlockToEVM(claim BurnClaim, signatureRel1 []byte, signatureRel2 []byte, signatureRel3 []byte) error {

	client, err := ethclient.Dial(os.Getenv("RPC_WS_URL"))
	if err != nil {
		return err
	}

	privateKeyHex := os.Getenv("RELAYER_EVM_PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return err
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 500000
	auth.GasPrice = gasPrice

	bridgeAddress := common.HexToAddress(os.Getenv("BRIDGE_ADDRESS"))

	bridge, err := abi.NewERC20Bridge(bridgeAddress, client)
	if err != nil {
		return err
	}

	claimTx := abi.ERC20BridgeClaim{
		EvmChainId: big.NewInt(claim.EvmChainId),
		Token:      claim.Token,
		To:         claim.Recipient,
		Amount:     claim.Amount,
		Nonce:      big.NewInt(claim.Nonce),
	}

	tx, err := bridge.Unlock(
		auth,
		claimTx,
		[][]byte{signatureRel1, signatureRel2, signatureRel3},
	)
	if err != nil {
		return err
	}

	log.Println("Unlock tx sent:", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed")
	}

	log.Println("Unlock confirmed")

	return nil
}
