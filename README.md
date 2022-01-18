# In-cluster kubectl

Kubectl is a fantastic tool. Typically users run it locally where it makes requests over the network
to a remote cluster. This is fine most of the time, but in some cases you may want kubectl to run
inside the cluster.

This project provides Docker images that contains kubectl with various versions
and a Go package to run arbitrary commands in the container that has kubectl.






