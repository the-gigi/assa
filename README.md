# In-cluster kubectl

Kubectl is a fantastic tool. Typically, users run it locally where it makes requests over the network
to a remote cluster. This is fine most of the time, but in some cases you may want kubectl to run
inside the cluster.

This project provides a Dockerfile to generate images that contain kubectl with various versions
and a Go package to run arbitrary commands in the container that has kubectl.

It can also push images to my Dockerhub account. To see all availabl images type:

```
$ curl -s "https://hub.docker.com/v2/repositories/g1g1/kubectl/tags?page_size=100" | jq '.results[].name' | sort

"1.21"
"1.22"
"1.23"
"1.24"
"1.25"
"1.26"
```

You can pull existing images without building by typing:

```
docker pull g1g1/kubectl:<tag> 
```


To build your own image for a particular version of kubectl change the VERSION at the top of the Makefile
and type `make`.

The `make push` command pushes the image to my Dockerhub account. If you want to push to your registry change the IMAGE variable in the Makefile. 













