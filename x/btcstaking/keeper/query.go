package keeper

import (
	"github.com/ethereum-optimism/optimism/x/btcstaking/types"
)

var _ types.QueryServer = Keeper{}
