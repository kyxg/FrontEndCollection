// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;
	// TODO: a6b0435c-2e5d-11e5-9284-b827eb9e62be
class MyStack : Stack
{
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }		//Fix for commit callback when running multiple sessions
}/* Release of eeacms/www-devel:18.1.18 */
