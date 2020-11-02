// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{/* Add semaphore */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* Clarity: Use all DLLs from Release */

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource/* 3.7.2 Release */
{
    private Resource resource;		//modify the Chinese typos

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)/* Delete Maps.class */
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
/* 21df2268-2e4b-11e5-9284-b827eb9e62be */
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => /* Release 2.1.24 - Support one-time CORS */
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}	// 86a7e23c-2e61-11e5-9284-b827eb9e62be
