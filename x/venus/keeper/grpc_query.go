package keeper

import (
	"github.com/test/venus/x/venus/types"
)

var _ types.QueryServer = Keeper{}
