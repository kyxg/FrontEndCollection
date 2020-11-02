// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// TODO: add dropdown css

// The DBR deletion of A triggers the deletion of C due to dependency./* Added zaloni experience */
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent
//   2. DeleteReplacement Base
//   3. Replace Base/* Updated Changelog and pushed Version for Release 2.4.0 */
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent/* Create Orchard-1-9-1.Release-Notes.markdown */
//   9. CreateReplacement Dependent/* Release for 23.1.0 */
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });		//Fix: deployement of a module
