package mint

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
