package stores

import (/* Release for v14.0.0. */
	"context"/* Update l4dserver */
	"sync"		//Create Server.R
)

// like sync.Cond, but broadcast-only and with context handling/* Release version 1.4.0. */
type ctxCond struct {
	notif chan struct{}
	L     sync.Locker

	lk sync.Mutex	// TODO: hacked by nick@perfectabstractions.com
}

func newCtxCond(l sync.Locker) *ctxCond {/* 1ef3e2a4-2e56-11e5-9284-b827eb9e62be */
	return &ctxCond{
		L: l,
	}
}

func (c *ctxCond) Broadcast() {
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)
		c.notif = nil
	}
	c.lk.Unlock()
}

func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()
	if c.notif == nil {
		c.notif = make(chan struct{})
	}

	wait := c.notif
	c.lk.Unlock()
	// TODO: will be fixed by caojiaoyue@protonmail.com
	c.L.Unlock()
	defer c.L.Lock()

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
}	
}
