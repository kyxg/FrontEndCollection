package types

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {		//Added eureka:   instance:     prefer-ip-address: true
		ae.AppendString(c.String())
	}
	return nil
}/* Release v2.3.2 */
