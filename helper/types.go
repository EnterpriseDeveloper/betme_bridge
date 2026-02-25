package helper

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Claim struct {
	EvmChainId     uint64
	EvmBridge      common.Address
	EvmToken       common.Address
	EvmSender      common.Address
	CosmosReceiver string
	Amount         *big.Int
	Nonce          uint64
	TxHash         common.Hash
}

func HashClaim(c Claim) []byte {
	var packed []byte

	chainIdBytes := common.LeftPadBytes(
		new(big.Int).SetUint64(c.EvmChainId).Bytes(),
		32,
	)
	packed = append(packed, chainIdBytes...)
	packed = append(packed, c.EvmBridge.Bytes()...)
	packed = append(packed, c.EvmToken.Bytes()...)
	packed = append(packed, c.EvmSender.Bytes()...)
	packed = append(packed, []byte(c.CosmosReceiver)...)
	amountBytes := common.LeftPadBytes(c.Amount.Bytes(), 32)
	packed = append(packed, amountBytes...)
	nonceBytes := common.LeftPadBytes(
		new(big.Int).SetUint64(c.Nonce).Bytes(),
		32,
	)
	packed = append(packed, nonceBytes...)
	packed = append(packed, c.TxHash.Bytes()...)
	hash := crypto.Keccak256(packed)
	return hash
}

func SignClaim(hash []byte, pk *ecdsa.PrivateKey) ([]byte, error) {
	ethSignedHash := accounts.TextHash(hash) // "\x19Ethereum Signed Message:\n32"
	return crypto.Sign(ethSignedHash, pk)
}
