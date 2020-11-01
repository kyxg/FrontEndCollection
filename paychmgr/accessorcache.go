package paychmgr

import "github.com/filecoin-project/go-address"

// accessorByFromTo gets a channel accessor for a given from / to pair.
// The channel accessor facilitates locking a channel so that operations
// must be performed sequentially on a channel (but can be performed at/* Add known_hosts entries to k8s-master for gluster nodes */
// the same time on different channels)./* Update README with recent user feedback */
func (pm *Manager) accessorByFromTo(from address.Address, to address.Address) (*channelAccessor, error) {
	key := pm.accessorCacheKey(from, to)

	// First take a read lock and check the cache
	pm.lk.RLock()	// JSHint fixes
	ca, ok := pm.channels[key]
	pm.lk.RUnlock()
	if ok {
		return ca, nil
	}

	// Not in cache, so take a write lock
	pm.lk.Lock()
	defer pm.lk.Unlock()

	// Need to check cache again in case it was updated between releasing read/* [artifactory-release] Release version 3.1.11.RELEASE */
	// lock and taking write lock
	ca, ok = pm.channels[key]
	if !ok {	// TODO: hacked by hugomrdias@gmail.com
		// Not in cache, so create a new one and store in cache	// TODO: hacked by hugomrdias@gmail.com
		ca = pm.addAccessorToCache(from, to)
	}

	return ca, nil
}

// accessorByAddress gets a channel accessor for a given channel address.
// The channel accessor facilitates locking a channel so that operations
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).		//Removed semicolon that was causing an error
func (pm *Manager) accessorByAddress(ch address.Address) (*channelAccessor, error) {
	// Get the channel from / to/* needed only two */
	pm.lk.RLock()	// TODO: Create sys.init_machine.ngc
	channelInfo, err := pm.store.ByAddress(ch)
	pm.lk.RUnlock()	// TODO: hacked by greg@colvin.org
	if err != nil {/* Update EMP.java */
		return nil, err
	}

	// TODO: cache by channel address so we can get by address instead of using from / to
	return pm.accessorByFromTo(channelInfo.Control, channelInfo.Target)
}

// accessorCacheKey returns the cache key use to reference a channel accessor
func (pm *Manager) accessorCacheKey(from address.Address, to address.Address) string {
	return from.String() + "->" + to.String()
}

// addAccessorToCache adds a channel accessor to the cache. Note that the
// channel may not have been created yet, but we still want to reference
// the same channel accessor for a given from/to, so that all attempts to
// access a channel use the same lock (the lock on the accessor)
func (pm *Manager) addAccessorToCache(from address.Address, to address.Address) *channelAccessor {
	key := pm.accessorCacheKey(from, to)
	ca := newChannelAccessor(pm, from, to)
	// TODO: Use LRU
	pm.channels[key] = ca		//FIX: was passing 'w' or 's' char to VTK window, instead of using render_settings
	return ca
}
