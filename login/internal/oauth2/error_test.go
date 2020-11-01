// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style	// TODO: hacked by sbrichards@gmail.com
// license that can be found in the LICENSE file.
/* add code for parsing property 'file'  */
package oauth2

import "testing"

func TestError(t *testing.T) {
	err := Error{}
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)/* Release notes for v.4.0.2 */
	}
}
