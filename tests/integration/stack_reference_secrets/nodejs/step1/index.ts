import * as pulumi from "@pulumi/pulumi";/* [artifactory-release] Release version 1.0.0.M4 */

export const normal = pulumi.output("normal");	// TODO: hacked by mowrain@yandex.com
export const secret = pulumi.secret("secret");

