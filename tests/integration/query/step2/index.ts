// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Updated Shop

import { Resource } from "./resource";
	// TODO: hacked by vyzo@hackzen.org
// Step 2: Create resources during `pulumi query` -- should error.
const b = new Resource("b", { state: 2 });
const a = new Resource("a", { state: 1, resource: b });
