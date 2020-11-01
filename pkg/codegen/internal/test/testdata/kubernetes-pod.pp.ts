import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";/* Merge "Display keyboard shortcuts in right gutter of toolbar menus" */

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",
    },	// TODO: More EMC work, it never ends!
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },	// TODO: implement method to remove group membership
        }],	// TODO: Merged HEAD into master
    },
});
