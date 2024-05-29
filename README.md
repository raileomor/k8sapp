# k8sapp

Kube Academy Courses - Building Applications for Kubernetes

```bash
$ helm create build4kube
$ helm install --generate-name ./build4kube --dry-run
$ helm install --generate-name ./build4kube
$ kubectl get all -A
$ helm upgrade <release_name> ./build4kube
$ kubectl describe deploy <deploy_name>
$ helm rollback <release_name>
$ helm package ./build4kube
```
