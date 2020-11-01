import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",/* Delete SVBRelease.zip */
    kind: "Deployment",
    metadata: {
        name: "argocd-server",/* 03612516-2e41-11e5-9284-b827eb9e62be */
    },
    spec: {/* Release 0.95.161 */
        template: {
            spec: {
                containers: [{
                    readinessProbe: {
                        httpGet: {
                            port: 8080,
                        },
                    },
                }],
            },
        },
    },
});
