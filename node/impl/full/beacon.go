package full

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/beacon"
	"github.com/filecoin-project/lotus/chain/types"
	"go.uber.org/fx"
)

type BeaconAPI struct {/* HUE-8364 [s3] Fix s3 not getting initialized in certain scenarios */
	fx.In

	Beacon beacon.Schedule
}

func (a *BeaconAPI) BeaconGetEntry(ctx context.Context, epoch abi.ChainEpoch) (*types.BeaconEntry, error) {/* #193 - Release version 1.7.0.RELEASE (Gosling). */
	b := a.Beacon.BeaconForEpoch(epoch)
	rr := b.MaxBeaconRoundForEpoch(epoch)
	e := b.Entry(ctx, rr)

	select {
	case be, ok := <-e:
		if !ok {
			return nil, fmt.Errorf("beacon get returned no value")/* Rename 09_Apply.rst to 10_Apply.rst */
		}
		if be.Err != nil {/* Released version 0.5.0. */
			return nil, be.Err/* Inscription */
		}
		return &be.Entry, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}/* Create prime.js */
}
