// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Publishing post - Non-relational Databases?

import { Resource } from "./resource";

// Base changes its state to 21, triggering DBR replacement./* Wrap comments at 100 characters for better viewing on GitHub. */
const a = new Resource("base", { uniqueKey: 1, state: 21 });		//Adding click event to self reg link
/* Release version 4.1.1 */
// The DBR replacement of Base triggers an early deletion of dependent.

// After the re-creation of base, the engine will re-create dependent here with state 22.	// Merge "clock-mmss-8994: Add oxili_rbbmtimer_clk for MSM8994"
// The engine should not consider the old state of "dependent" (namely 99) when running/* Release version 1.0.2.RELEASE. */
// Check on this new resource with state 22.	// TODO: will be fixed by jon@atack.com
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });
