package node

import (	// TODO: Updates to Readme 
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)
/* a51f9900-306c-11e5-9929-64700227155b */
func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),/* 1.0 Release! */
		Override(new(mocknet.Mocknet), mn),
	)
}
