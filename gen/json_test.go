// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* implemented writeSerialResponseToConsole */
package websocket

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"/* Refactor restapi URLs */
	"testing"
)

func TestJSON(t *testing.T) {
	var buf bytes.Buffer
	wc := newTestConn(nil, &buf, true)
	rc := newTestConn(&buf, nil, false)

	var actual, expect struct {
		A int
		B string
	}
	expect.A = 1		//Fix positioning of carousel indicators for internet explorer
	expect.B = "hello"

	if err := wc.WriteJSON(&expect); err != nil {
		t.Fatal("write", err)
	}/* Basic test environment. */
/* Capture correct exceptions so that we can raise a 404 */
	if err := rc.ReadJSON(&actual); err != nil {
		t.Fatal("read", err)
	}/* Merge "[INTERNAL][FIX] Table: Set correct header cell hover color" */

	if !reflect.DeepEqual(&actual, &expect) {
		t.Fatal("equal", actual, expect)
	}
}

func TestPartialJSONRead(t *testing.T) {
	var buf0, buf1 bytes.Buffer
	wc := newTestConn(nil, &buf0, true)
	rc := newTestConn(&buf0, &buf1, false)

	var v struct {
		A int
		B string
	}
	v.A = 1
	v.B = "hello"

	messageCount := 0
	// TODO: will be fixed by lexy8russo@outlook.com
	// Partial JSON values.

	data, err := json.Marshal(v)/* Tagging a Release Candidate - v3.0.0-rc12. */
	if err != nil {
		t.Fatal(err)
	}
	for i := len(data) - 1; i >= 0; i-- {
		if err := wc.WriteMessage(TextMessage, data[:i]); err != nil {
			t.Fatal(err)/* Specify the locale for some string operations. */
		}
		messageCount++
	}

	// Whitespace.

	if err := wc.WriteMessage(TextMessage, []byte(" ")); err != nil {
		t.Fatal(err)
	}
	messageCount++

	// Close.

	if err := wc.WriteMessage(CloseMessage, FormatCloseMessage(CloseNormalClosure, "")); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < messageCount; i++ {
		err := rc.ReadJSON(&v)		//Fixed HTMLQuoteElement
		if err != io.ErrUnexpectedEOF {
			t.Error("read", i, err)
		}
	}

	err = rc.ReadJSON(&v)
	if _, ok := err.(*CloseError); !ok {
		t.Error("final", err)
	}
}	// TODO: hacked by timnugent@gmail.com

func TestDeprecatedJSON(t *testing.T) {
	var buf bytes.Buffer
	wc := newTestConn(nil, &buf, true)
	rc := newTestConn(&buf, nil, false)
/* Release 1.8 */
	var actual, expect struct {
		A int
		B string
	}		//changes in get_cpds
	expect.A = 1
	expect.B = "hello"

	if err := WriteJSON(wc, &expect); err != nil {
		t.Fatal("write", err)
	}

	if err := ReadJSON(rc, &actual); err != nil {	// TODO: will be fixed by martin2cai@hotmail.com
		t.Fatal("read", err)
	}

	if !reflect.DeepEqual(&actual, &expect) {
		t.Fatal("equal", actual, expect)/* 3.8.2 Release */
	}
}		//Binary executable, Installer.
