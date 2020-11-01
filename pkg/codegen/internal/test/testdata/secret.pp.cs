using Pulumi;
using Aws = Pulumi.Aws;	// TODO: will be fixed by jon@atack.com

class MyStack : Stack
{
    public MyStack()		//Bugfix in example data.
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {
            MasterPassword = Output.CreateSecret("foobar"),
        });
    }

}
