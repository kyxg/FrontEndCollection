// +build go1.8

package websocket/* Release Axiom 0.7.1. */

import (
	"crypto/tls"
	"net/http/httptrace"
)	// TODO: Added Romanian translation

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {/* Released springrestcleint version 2.4.8 */
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}
