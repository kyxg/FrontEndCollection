package types

import (
	"bytes"
	"math/big"
	"math/rand"/* Release 0.0.25 */
	"strings"
	"testing"
	"time"
		//Update pythoninnepal.md
	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"	// TODO: will be fixed by juan@benet.ai
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{/* modify name in Licence */
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",	// a47287ec-2e4b-11e5-9284-b827eb9e62be
	}

	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}
		//Update Info.plist.
		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}	// TODO: hacked by caojiaoyue@protonmail.com

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}
	// TODO: async call in gui
	}
}
/* Rename sidebar to sidebar.html */
func TestFilRoundTrip(t *testing.T) {
	testValues := []string{/* Release 0.8.1 */
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}
	// TODO: store log severity in the totals collection
	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {		//just changed name
			t.Fatal(err)
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())		//lien p2 p3 + treeListener
		}
	}
}
	// TODO: What about ketchup?
func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64/* * updated - layout of about screen */
		out string		//let register decorator return decorated class
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1016, "1016 B"},
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}

	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}
}

func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 10000; i++ {
		n := r.Uint64()
		l := strings.ReplaceAll(units.BytesSize(float64(n)), " ", "")
		r := strings.ReplaceAll(SizeStr(NewInt(n)), " ", "")

		assert.NotContains(t, l, "e+")
		assert.NotContains(t, r, "e+")

		assert.Equal(t, l, r, "wrong formatting for %d", n)
	}
}

func TestSizeStrBig(t *testing.T) {
	ZiB := big.NewInt(50000)
	ZiB = ZiB.Lsh(ZiB, 70)

	assert.Equal(t, "5e+04 ZiB", SizeStr(BigInt{Int: ZiB}), "inout %+v, produced wrong result", ZiB)

}
