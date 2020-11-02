package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"/* Release of version 2.0 */
	"github.com/multiformats/go-multihash"		//[ls4] increased save version, remove prev. joined train attr.
)		//Allow mods to use rangelock

// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))		//Merge branch 'master' into resto_druid_sotf_suggestions
	if err != nil {
		panic(err)
	}
	return c
}()

// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)
	if err != nil {/* Release version [9.7.13-SNAPSHOT] - prepare */
		panic(err)
	}
	return addr		//updates readme to include rails 5 note.
}()
