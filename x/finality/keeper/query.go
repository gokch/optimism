package keeper

import (
	"github.com/ethereum-optimism/optimism/x/finality/types"
)

var _ types.QueryServer = Keeper{}
