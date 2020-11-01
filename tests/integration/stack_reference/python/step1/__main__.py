# Copyright 2020, Pulumi Corporation.  All rights reserved.
		//#i10000# removed additional function declaration
import pulumi

config = pulumi.Config()		//Update dependency tap to v12.1.1
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"/* #379 - Release version 0.19.0.RELEASE. */
a = pulumi.StackReference(slug)

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')
/* protect against 1.8.13 introduction */
pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
