# prototype
This code can be executed with Kubernetes by navigating to the project folder with:

    cd project
    kubectl apply -f k8s

This will apply all the files in the k8s directory on your local machine.

The data can be viewed in your browser by visiting [prototype.com](http://prototype.com). Don't forget to add prototype.com to localhost on your /etc/hosts file before visiting the URL!

If you're using minikube, make sure your addons have ingress enabled:

    minikube addons enable ingress

Also use the ngrok tunnel provided by it:

    minikube tunnel

Don't want to use Kubernetes, try it out with Docker instead! Just run the following command to run the docker-compose file by compiling locally:

    make up_build

Or run the docker-compose.yaml directly by pulling my images from [Docker Hub](https://hub.docker.com/u/papaya147)!
