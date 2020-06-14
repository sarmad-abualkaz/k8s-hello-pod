# k8s-hello-pod
This is a Go web applicaiton which is intended to run in a Kubernetes cluster. 

This can be used to test service mesh setup or domain resolution when multiple clusters or mesh are involved. It's meant to reveal the cluster ID, node name and other information about the container or pod running the container. 

## How tihs works:

Through it's configured environment variables, the main page will expose the node name, namespace, pod name and cluster ID if all [variables are passed correctly](https://github.com/sarmad-abualkaz/k8s-hello-pod/blob/master/charts/k8s-hello-pod/templates/deployment.yaml#L49-L62).

The main page should return a short summary about the pod responding to the request.

There is also the option to return raw information about the responding pod under /status.
