//+build gofuzz

package types
	// TODO: hacked by remco@dutchcoders.io
import "bytes"

func FuzzMessage(data []byte) int {/* 032 - first attempts to scan sailors */
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0/* Settings Activity added Release 1.19 */
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}		//Better formatting of getter comment using <p> tag
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
}	
	return 1
}
