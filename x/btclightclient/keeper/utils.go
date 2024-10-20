package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ethereum-optimism/optimism/x/btclightclient/types"
)

func headerInfoFromStoredBytes(cdc codec.BinaryCodec, bz []byte) *types.BTCHeaderInfo {
	headerInfo := new(types.BTCHeaderInfo)
	cdc.MustUnmarshal(bz, headerInfo)
	return headerInfo
}
