getpod
======

Address pods quickly without using the clipboard or listing all pods first. Here's a typical use case:

```
$ kubectl logs `getpod kube-proxy` -f
```

What have we gained? This line replaces two common approaches. In many cases we would use the clipboard or type in the hash component of the name:

```
$ kubectl get po | grep kube-proxy
kube-proxy-cf2df                            1/1     Running   0          5h
kube-proxy-8dj6v                            1/1     Running   0          5h
kube-proxy-s6wvq                            1/1     Running   0          5h
$ kubectl logs kube-proxy-cf2df -f
```

Why not just script it? Be my guest:

```
$ kubectl logs `kubectl get po --no-headers | grep kube-proxy | cut -d' ' -f1 | head -n1` -f
```

Why not just a shell script? Because it may not work on Windows. It is hard to overestimate the number of Kubernetes users using `kubectl` via Git Bash or Power Shell.

## Usage
```
getpod [-kubeconfig PATH] POD
```

## Build
```
make
```