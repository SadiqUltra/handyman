# Handyman

### Tech
* GoLang
* GraphQL
* Postgresql
* Docker
* pgadmin
* kubernetes

### Installation

Handyman requires `Docker`,`Docker-compose` to run.

Install the dependencies and devDependencies and start the server.

```sh
$ docker-compose up
```

Open [this address](http://localhost:8080/) in your preferred browser.
```sh
http://localhost:8080/
```

### For Kubernetes cluster
For local environment we will use `minikube`. At first make sure its up and running.
```sh
$ minikube start
```

Run those following command in sequence
```sh
$ kubectl create -f postgres-secret.yaml
$ kubectl apply -f postgres-db-pv.yaml
$ kubectl apply -f postgres-db-pvc.yaml
$ kubectl apply -f postgres-db-deployment.yaml
$ kubectl apply -f postgres-db-service.yaml
$ kubectl apply -f app-postgres-deployment.yaml
$ kubectl apply -f app-postgres-service.yaml
```

Get the exposed url and port

```sh 
$ minikube service handyman-app-postgres --url
```
Url will be look like this one: `http://192.168.xxx.xxx:xxxxx`
