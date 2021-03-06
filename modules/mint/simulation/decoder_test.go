package simulation

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmkv "github.com/tendermint/tendermint/libs/kv"

	"github.com/irisnet/irishub/modules/mint/types"
	"github.com/irisnet/irishub/modules/random/simulation"
	"github.com/irisnet/irishub/simapp"
)

func TestDecodeStore(t *testing.T) {
	minter := types.NewMinter(time.Now().UTC(), sdk.NewIntWithDecimal(2, 9))
	cdc, _ := simapp.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	kvPairs := tmkv.Pairs{
		tmkv.Pair{Key: types.MinterKey, Value: cdc.MustMarshalBinaryBare(&minter)},
		tmkv.Pair{Key: []byte{0x99}, Value: []byte{0x99}},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Minter", fmt.Sprintf("%v\n%v", minter, minter)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
