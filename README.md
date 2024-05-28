# k8sapp

Kube Academy Courses - Building Applications for Kubernetes

With `kustomize`:

```bash
$ kustomize build base
$ kustomize build overlays/production
$ kustomize build overlays/production | kubectl apply -f -
```

With `kubectl`:

```bash
$ kubectl kustomize base
$ kubectl kustomize overlays/production
$ kubectl apply -k overlays/production
$ kubectl delete -k overlays/production
```
