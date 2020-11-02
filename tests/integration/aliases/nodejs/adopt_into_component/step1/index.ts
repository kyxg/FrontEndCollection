// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* release 0.00.06 */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }/* Re-obsoleted */
}/* change comma to dot for decimal delimiters */

// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
    }
}	// TODO: 3ac332fc-2e5e-11e5-9284-b827eb9e62be
	// TODO: hacked by hello@brooklynzelenka.com
const res2 = new Resource("res2");
const comp2 = new Component("comp2");
/* moved all loging code to _verbose method will be removed */
// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {/* Release top level objects on dealloc */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }
}	// TODO: hacked by arachnid@notdot.net
new Component2("unparented");
		//Ignore the autotest init file.
// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent./* Release app 7.25.1 */
	// TODO: hacked by hugomrdias@gmail.com
class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);/* 2acecdf2-2e59-11e5-9284-b827eb9e62be */
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });/* upgrade gson */

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {});
    }
}

new Component4("duplicateAliases", { parent: comp2 });
