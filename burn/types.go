package burn

import (
	"crypto/ecdsa"
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

	chainID := new(big.Int).SetInt64(claim.EvmChainId)

	packed = append(packed, common.LeftPadBytes(chainID.Bytes(), 32)...)
	packed = append(packed, claim.Bridge.Bytes()...)
	packed = append(packed, claim.Token.Bytes()...)
	packed = append(packed, claim.Recipient.Bytes()...)
	packed = append(packed, common.LeftPadBytes(claim.Amount.Bytes(), 32)...)

	nonce := new(big.Int).SetInt64(claim.Nonce)
	packed = append(packed, common.LeftPadBytes(nonce.Bytes(), 32)...)

	return crypto.Keccak256(packed)
}

func SignClaim(hash []byte, pk *ecdsa.PrivateKey) ([]byte, error) {

	ethHash := EthereumSignedHash(hash)

	signature, err := crypto.Sign(ethHash, pk)
	if err != nil {
		return nil, err
	}

	// convert v from 0/1 → 27/28
	signature[64] += 27

	return signature, nil
}

func EthereumSignedHash(hash []byte) []byte {
	prefix := []byte("\x19Ethereum Signed Message:\n32")
	return crypto.Keccak256(append(prefix, hash...))
}
