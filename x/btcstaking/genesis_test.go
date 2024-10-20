package btcstaking_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ethereum-optimism/optimism/x/testutil/keeper"
	"github.com/ethereum-optimism/optimism/x/testutil/nullify"
	"github.com/ethereum-optimism/optimism/x/btcstaking"
	"github.com/ethereum-optimism/optimism/x/btcstaking/types"
)

func TestGenesis(t *testing.T) {
	p := types.DefaultParams()
	genesisState := types.GenesisState{
		Params: []*types.Params{&p},
	}

	k, ctx := keepertest.BTCStakingKeeper(t, nil, nil, nil, nil)
	btcstaking.InitGenesis(ctx, *k, genesisState)
	got := btcstaking.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
