package types

import (
	"context"

	bbn "github.com/babylonlabs-io/babylon/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	btcctypes "github.com/ethereum-optimism/optimism/x/btccheckpoint/types"
	btclctypes "github.com/ethereum-optimism/optimism/x/btclightclient/types"
)

type BTCLightClientKeeper interface {
	GetBaseBTCHeader(ctx context.Context) *btclctypes.BTCHeaderInfo
	GetTipInfo(ctx context.Context) *btclctypes.BTCHeaderInfo
	GetHeaderByHash(ctx context.Context, hash *bbn.BTCHeaderHashBytes) *btclctypes.BTCHeaderInfo
}

type BtcCheckpointKeeper interface {
	GetParams(ctx context.Context) (p btcctypes.Params)
}

type FinalityKeeper interface {
	HasTimestampedPubRand(ctx context.Context, fpBtcPK *bbn.BIP340PubKey, height uint64) bool
}

type BtcStakingHooks interface {
	AfterFinalityProviderActivated(ctx context.Context, fpPk *bbn.BIP340PubKey) error
}

type IncentiveKeeper interface {
	IndexRefundableMsg(ctx context.Context, msg sdk.Msg)
}
