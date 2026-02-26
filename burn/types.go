package burn

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type BurnClaim struct {
	EvmChainId int64
	Bridge     common.Address
	Token      common.Address
	Recipient  common.Address
	Amount     *big.Int
	Nonce      int64
}

func ParseBurnEvent(event abci.Event) BurnClaim {

	attrs := map[string]string{}

	for _, attr := range event.Attributes {
		attrs[string(attr.Key)] = string(attr.Value)
	}

	chainID, _ := strconv.ParseInt(attrs["chain_id"], 10, 64)
	nonce, _ := strconv.ParseInt(attrs["nonce"], 10, 64)

	amount := new(big.Int)
	amount.SetString(attrs["amount"], 10)

	return BurnClaim{
		EvmChainId: chainID,
		Bridge:     common.HexToAddress(attrs["bridge"]),
		Token:      common.HexToAddress(attrs["token"]),
		Recipient:  common.HexToAddress(attrs["recipient"]),
		Amount:     amount,
		Nonce:      nonce,
	}
}

func HashBurnClaim(claim BurnClaim) []byte {

	var packed []byte

	// chainId (uint256)
	chainID := new(big.Int).SetInt64(claim.EvmChainId)
	packed = append(packed, common.LeftPadBytes(chainID.Bytes(), 32)...)

	// bridge (address)
	packed = append(packed, claim.Bridge.Bytes()...)

	// token (address)
	packed = append(packed, claim.Token.Bytes()...)

	// recipient (address)
	packed = append(packed, claim.Recipient.Bytes()...)

	// amount (uint256)
	packed = append(packed, common.LeftPadBytes(claim.Amount.Bytes(), 32)...)

	// nonce (uint256)
	nonce := new(big.Int).SetInt64(claim.Nonce)
	packed = append(packed, common.LeftPadBytes(nonce.Bytes(), 32)...)

	return crypto.Keccak256(packed)
}

func SignClaim(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {

	ethHash := EthereumSignedHash(hash)

	signature, err := crypto.Sign(ethHash, privateKey)
	if err != nil {
		return nil, err
	}

	// fix V value (27/28)
	signature[64] += 27

	return signature, nil
}

func EthereumSignedHash(hash []byte) []byte {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(hash))
	return crypto.Keccak256([]byte(prefix), hash)
}
