package burn

import (
	abi "bridge_betme/abi"
	"math/big"
	"strconv"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/ethereum/go-ethereum/common"
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
