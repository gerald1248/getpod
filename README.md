getpod
======

Access pods quickly without using the clipboard or listing all pods first. Here's a typical use case:

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

What about a one-liner? It's straightforward but hardly concise:

```
$ kubectl logs `kubectl get po --no-headers | grep kube-proxy | cut -d' ' -f1 | head -n1` -f
```

Why not just a shell script? Mainly because it probably will not work as intended on Windows. It is hard to overestimate the number of Kubernetes users using `kubectl` via Git Bash or PowerShell.

## Usage
```
$ getpod -h
Usage: getpod [-kubeconfig=PATH] [-a] [-n NAMESPACE] REGEX
  -a	return all matching pods
  -kubeconfig string
    	(optional) absolute path to the kubeconfig file
  -n string
    	namespace
```

The search expression is interpreted by the Golang `regexp` package (sadly not PCRE).

## Run
To try `gopod` on your computer, download one of the release binaries above (Linux, Mac, Windows).

## Build
```
$ make
$ sudo make install
```

