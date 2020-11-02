// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager

import (
	"io/ioutil"
/* test: add RawMessageQueue test module to compilation */
	"github.com/sirupsen/logrus"
)
/* I FIXED THE SHIT WITH THE STUFF */
func init() {
	logrus.SetOutput(ioutil.Discard)
}
