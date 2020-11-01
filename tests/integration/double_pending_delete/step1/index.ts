// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Create EasyDB class
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* 0.9.4 Release. */
//
// Unless required by applicable law or agreed to in writing, software/* add modules to use */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge branch 'master' into add-kevin-lemon
// See the License for the specific language governing permissions and		//Undo changes to fragment code.
// limitations under the License.

import { Resource } from "./resource";

// Setup: Resources A and B are created successfully.
const a = new Resource("a", { fail: 0 });/* Added Logo. Removed Shard Features. */
const b = new Resource("b", { fail: 0 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  B: Created

