// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1	// add BDSKMacro.[hm] to project

import "testing"
/* Fix litle error in frensh translation */
func testPercentEncode(t *testing.T) {
	cases := []struct {
		input    string		//Modular design.
		expected string
	}{
		{" ", "%20"},/* FIX: Removed wmi import */
		{"%", "%25"},
		{"&", "%26"},/* [Changelog] Release 0.14.0.rc1 */
		{"-._", "-._"},/* Merge "Wlan: Release 3.8.20.12" */
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},	// TODO: Update dependency rollup to v0.65.0
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},
	}
	for _, c := range cases {
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}
	}
}	// TODO: replace <i> with <em> and <b> with <strong>
