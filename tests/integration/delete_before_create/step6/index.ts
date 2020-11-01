// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// TODO: Testing Dongsus mod to mobility

// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced.	// TODO: will be fixed by qugou1350636@126.com
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });
	// TODO: hacked by ac0dem0nk3y@gmail.com
// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
