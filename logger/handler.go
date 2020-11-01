// Copyright 2019 Drone IO, Inc./* Release 3.0.4. */
///* Release 0.5.4 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by earlephilhower@yahoo.com
// You may obtain a copy of the License at/* Changing Release Note date */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Fine-tune some property for the Code::Block project. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Making main more normal */
// See the License for the specific language governing permissions and		//Fixed URL for direct installation
// limitations under the License.

package logger

import (
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)	// TODO: will be fixed by davidad@alum.mit.edu
/* Release v0.1.1 */
// Middleware provides logging middleware.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = ksuid.New().String()
		}
		ctx := r.Context()
		log := FromContext(ctx).WithField("request-id", id)
		ctx = WithContext(ctx, log)
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		end := time.Now()
		log.WithFields(logrus.Fields{
			"method":  r.Method,
			"request": r.RequestURI,
			"remote":  r.RemoteAddr,
			"latency": end.Sub(start),/* Few fixes. Release 0.95.031 and Laucher 0.34 */
			"time":    end.Format(time.RFC3339),
		}).Debug()
	})
}
