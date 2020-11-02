// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved./* make modular ns be comparable */
// Use of this source code is governed by a BSD-style/* Fixed typo, thx https://news.ycombinator.com/item?id=6411235. */
// license that can be found in the LICENSE file.
		//Create qualimap.sh
// +build go1.8
/* Update scrolling for RHS threads (#2803) */
package websocket

import "crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {		//Create Cuisine.js
		return &tls.Config{}
	}
	return cfg.Clone()
}
