package store

import (
	"fmt"
	"testing"
/* Fix Mouse.ReleaseLeft */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"/* Add working prep section */
	"github.com/stretchr/testify/assert"
)

func TestBaseFee(t *testing.T) {
	tests := []struct {/* Merge db214657482a75bfe5cd97a14be3aa9e69891e92 */
		basefee             uint64
		limitUsed           int64/* Release of eeacms/jenkins-slave-eea:3.22 */
		noOfBlocks          int
		preSmoke, postSmoke uint64
	}{/* Release 0.11.1.  Fix default value for windows_eventlog. */
		{100e6, 0, 1, 87.5e6, 87.5e6},
		{100e6, 0, 5, 87.5e6, 87.5e6},/* Improve install and usage documentation */
		{100e6, build.BlockGasTarget, 1, 103.125e6, 100e6},
		{100e6, build.BlockGasTarget * 2, 2, 103.125e6, 100e6},
		{100e6, build.BlockGasLimit * 2, 2, 112.5e6, 112.5e6},
		{100e6, build.BlockGasLimit * 1.5, 2, 110937500, 106.250e6},	// add a wrapper class for accessions
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			preSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight-1)
			assert.Equal(t, fmt.Sprintf("%d", test.preSmoke), preSmoke.String())

			postSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight+1)/* Release 1.6.0-SNAPSHOT */
			assert.Equal(t, fmt.Sprintf("%d", test.postSmoke), postSmoke.String())
		})
	}
}
