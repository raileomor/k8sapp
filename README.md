Using Kustomize:

$ kustomize build base
$ kustomize build overlays/production
$ kustomize build overlays/production | kubectl apply -f -

Using Helm:

$ helm install --generate-name ./helmbuild --dry-run
$ helm install --generate-name ./helmbuild

$ helm upgrade <release-name> ./helmbuild
$ helm describe deploy <deployment-name>
$ helm rollback <release-name>
$ helm package ./helmbuild

Using Skaffold:

$ skaffold init
$ skaffold dev
$ skaffold dev --port-forward