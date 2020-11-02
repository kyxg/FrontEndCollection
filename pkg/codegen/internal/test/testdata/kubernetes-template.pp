resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"/* disabled buffer overflow checks for Release build */
	kind = "Deployment"
	metadata = {/* Add some Release Notes for upcoming version */
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [
					{
						readinessProbe = {/* AG: cf system spec uses route53 outfile */
							httpGet = {
								port = 8080
							}
						}
					}
				]
			}
		}
	}
}
