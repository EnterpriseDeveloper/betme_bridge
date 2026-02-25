package helper

import sdk "github.com/cosmos/cosmos-sdk/types"

func Init() {
	config := sdk.GetConfig()

	config.SetBech32PrefixForAccount("bettery", "betterypub")
	config.SetBech32PrefixForValidator("betteryvaloper", "betteryvaloperpub")
	config.SetBech32PrefixForConsensusNode("betteryvalcons", "betteryvalconspub")

	config.Seal()
}
