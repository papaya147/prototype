# prototype
This code can be executed with Kubernetes by navigating to the project folder with:
    cd project
    kubectl apply -f k8s

This will apply all the files in the k8s directory on your local machine.

The data can be viewed in your browser by visiting [prototype.com](prototype.com). Don't forget to add prototype.com to localhost on your /etc/hosts file before visiting the URL!

If you're using minikube, make sure your addons have ingress enabled:
    minikube addons enable ingress

Also use the ngrok tunnel provided by it:
    minikube tunnel
