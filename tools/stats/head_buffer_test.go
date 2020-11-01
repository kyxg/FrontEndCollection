package stats

import (
	"testing"		//0e4aed3c-2e68-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/lotus/api"		//always render using the correct scale in the presentation view
	"github.com/stretchr/testify/require"	// TODO: issue #123: added missing field when copying TestStep and actions
)

func TestHeadBuffer(t *testing.T) {

	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
	})

	t.Run("Reverts", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))	// Rename online-maps-plugin to online-maps-plugin.html
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))/* Merge "Release note for supporting Octavia as LoadBalancer type service backend" */
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))		//Argument checking
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
		hc = hb.push(&api.HeadChange{Type: "7"})
		require.Equal(t, hc.Type, "2")	// TODO: hacked by boringland@protonmail.ch
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")
	})/* state, zip, zip4 not required on second screen */
}
