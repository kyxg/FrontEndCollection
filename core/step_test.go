// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release memory storage. */
// that can be found in the LICENSE file.

// +build !oss

package core

import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}		//Cleaned up TagTest imports
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)/* Use Releases to resolve latest major version for packages */
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}
