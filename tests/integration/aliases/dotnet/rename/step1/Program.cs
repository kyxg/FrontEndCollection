// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by igor@soramitsu.co.jp
		//Merge branch 'master' into new-contrib
using System.Threading.Tasks;
using Pulumi;		//[FIX] portal: error during portal creation

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

class Program
{
    static Task<int> Main(string[] args)
{    
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");
        });/* Create DateTools.ahk */
    }
}
