// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* 2b2d945c-2e65-11e5-9284-b827eb9e62be */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db
		//Actually center the loading icon. Suggested by Keri Henare.
type nopLocker struct{}
	// TODO: hacked by brosner@gmail.com
func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}/* - updated dev status. */
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
