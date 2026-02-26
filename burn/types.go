package burn

import (
	abi "bridge_betme/abi"
	"crypto/ecdsa"
	"math/big"
	"strconv"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func ParseBurnEvent(event abci.Event) abi.ERC20BridgeClaim {

	attrs := map[string]string{}

	for _, attr := range event.Attributes {
		attrs[string(attr.Key)] = string(attr.Value)
	}

	chainID, _ := strconv.ParseInt(attrs["chain_id"], 10, 64)
	nonce, _ := strconv.ParseInt(attrs["nonce"], 10, 64)

	amount := new(big.Int)
	amount.SetString(attrs["amount"], 10)

	return abi.ERC20BridgeClaim{
		EvmChainId: big.NewInt(chainID),
		Token:      common.HexToAddress(attrs["token"]),
		To:         common.HexToAddress(attrs["recipient"]),
		Amount:     amount,
		Nonce:      big.NewInt(nonce),
	}
}

func HashBurnClaim(claim abi.ERC20BridgeClaim) []byte {

	var packed []byte
	packed = append(packed, common.LeftPadBytes(claim.EvmChainId.Bytes(), 32)...)
	packed = append(packed, claim.Token.Bytes()...)
	packed = append(packed, claim.To.Bytes()...)
	packed = append(packed, common.LeftPadBytes(claim.Amount.Bytes(), 32)...)
	packed = append(packed, common.LeftPadBytes(claim.Nonce.Bytes(), 32)...)

	return crypto.Keccak256(packed)
}

func SignClaim(hash []byte, pk *ecdsa.PrivateKey) ([]byte, error) {
	prefixedHash := crypto.Keccak256(
		[]byte("\x19Ethereum Signed Message:\n32"),
		hash,
	)

	signature, err := crypto.Sign(prefixedHash, pk)
	if err != nil {
		return nil, err
	}

	// convert v from 0/1 → 27/28
	signature[64] += 27

	return signature, nil
}
