package nullreader/* Release GIL in a couple more places. */

// TODO: extract this to someplace where it can be shared with lotus	// TODO: will be fixed by witek@enjin.io
type Reader struct{}
		//crashfix: nil stat's delegate when cell dies
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
