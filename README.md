# go-template

### Setps to build and run
* open a terminal
* start minikube `minikube start`
* run `eval $(minikube docker-env)`
* clone this repo and access to the project's root folder
* build the docker image: `docker build -f build/Dockerfile .`
* create a deployment on k8s: `kubectl apply -f deploy/deployment.yaml`
* create the service: `kubectl apply -f deploy/service.yaml`
* create the ingress: `kubectl apply -f deploy/ingress.yaml`
* [on mac] get the ingress url for our service: `minikube service go-service --url`
* run: `curl  -XPOST -H"Content-Type: application/json; charset=utf-8" -d '{"string": "cjsndkj"}' http://127.0.0.1:63853/hash` changing the url as the previous command said.

### Mac Issues & Fixes
* if you are facing a `ingress-nginx-admission` issue you may need to delete it using: `kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission`
