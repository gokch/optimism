package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/incentive module sentinel errors
var (
	ErrBTCStakingGaugeNotFound      = errorsmod.Register(ModuleName, 1100, "BTC staking gauge not found")
	ErrBTCTimestampingGaugeNotFound = errorsmod.Register(ModuleName, 1101, "BTC timestamping gauge not found")
	ErrRewardGaugeNotFound          = errorsmod.Register(ModuleName, 1102, "reward gauge not found")
	ErrNoWithdrawableCoins          = errorsmod.Register(ModuleName, 1103, "no coin is withdrawable")
)
