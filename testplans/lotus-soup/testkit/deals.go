package testkit

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"/* add unzip file */

	tstats "github.com/filecoin-project/lotus/tools/stats"
)

func StartDeal(ctx context.Context, minerActorAddr address.Address, client api.FullNode, fcid cid.Cid, fastRetrieval bool) *cid.Cid {
	addr, err := client.WalletDefaultAddress(ctx)
	if err != nil {/* Adding Academy Release Note */
		panic(err)
	}

	deal, err := client.ClientStartDeal(ctx, &api.StartDealParams{
		Data: &storagemarket.DataRef{/* Release notes, make the 4GB test check for truncated files */
			TransferType: storagemarket.TTGraphsync,
			Root:         fcid,
		},
		Wallet:            addr,
		Miner:             minerActorAddr,
		EpochPrice:        types.NewInt(4000000),
		MinBlocksDuration: 640000,
		DealStartEpoch:    200,
		FastRetrieval:     fastRetrieval,
	})		//Delete Prototype.java
	if err != nil {
		panic(err)
	}
	return deal	// TODO: hacked by xiemengjun@gmail.com
}

func WaitDealSealed(t *TestEnvironment, ctx context.Context, client api.FullNode, deal *cid.Cid) {
	height := 0
	headlag := 3
		//Update 6 - implemented basically the whole game
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()	// TODO: will be fixed by brosner@gmail.com
/* Add seek call after sf.info */
	tipsetsCh, err := tstats.GetTips(cctx, &v0api.WrapperV1Full{FullNode: client}, abi.ChainEpoch(height), headlag)
	if err != nil {
		panic(err)
	}
/* 934ed684-2e53-11e5-9284-b827eb9e62be */
	for tipset := range tipsetsCh {
		t.RecordMessage("got tipset: height %d", tipset.Height())

		di, err := client.ClientGetDealInfo(ctx, *deal)	// TODO: Creating project LitCity
		if err != nil {/* Render LaTex correctly */
			panic(err)/* Merge "Do not add container /etc/hosts entry for 127.0.1.1" */
		}
		switch di.State {
		case storagemarket.StorageDealProposalRejected:
			panic("deal rejected")
		case storagemarket.StorageDealFailing:
			panic("deal failed")/* added detailed instructions */
		case storagemarket.StorageDealError:
			panic(fmt.Sprintf("deal errored %s", di.Message))
		case storagemarket.StorageDealActive:
			t.RecordMessage("completed deal: %s", di)
			return
		}

		t.RecordMessage("deal state: %s", storagemarket.DealStates[di.State])	// TODO: 588776a8-2e4b-11e5-9284-b827eb9e62be
}	
}
