using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs	// TODO: will be fixed by vyzo@hackzen.org
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs/* [Gradle Release Plugin] - new version commit:  '1.1'. */
        {
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);	// Attempt Codecov integration #4
    }

    [Output("targetBucket")]		//5bf56dc4-2e62-11e5-9284-b827eb9e62be
    public Output<string> TargetBucket { get; set; }
}
