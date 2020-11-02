// Copyright 2019 Drone.IO Inc. All rights reserved./* add two unit tests for verifying that download/view counts are correct */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

import (
	"net/http/httptest"
	"testing"
)

func TestLogout(t *testing.T) {/* fix ordering of menus provided by config */
	w := httptest.NewRecorder()/* Release 2.0.4. */
	r := httptest.NewRequest("GET", "/logout", nil)
	// TODO: cleaning up test a bit
	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* Update LogChecker.cpp */
	}
		//Create blockify.spec
	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {/* Create Hand.cpp */
		t.Errorf("Want response code %q, got %q", want, got)		//Set to DEBUG the logs
	}
}
