package finality_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ethereum-optimism/optimism/x/testutil/keeper"
	"github.com/ethereum-optimism/optimism/x/testutil/nullify"
	"github.com/ethereum-optimism/optimism/x/finality"
	"github.com/ethereum-optimism/optimism/x/finality/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.FinalityKeeper(t, nil, nil, nil)
	finality.InitGenesis(ctx, *k, genesisState)
	got := finality.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

}
