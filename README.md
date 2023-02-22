# URL Shortener

## Stack
- Go
- Redis
- Docker
- Kompose
- minikube

## Running app

### Using docker / docker-compose
With building local app image
```
docker-compose -f docker-compose.dev.yml up --build
```
Or using remote image
```
docker-compose up
```

### Using minikube / kubectl
```
kubectl apply -f k8s/
```

## Useful commands
Generating k8s files, k8s configutaion files `k8s/*.yaml`  are generated using **Kompose** tool based on `docker-compose.yml`
```
kompose convert -o k8s/
```

apply config 
```
kubectl apply -f k8s/
```

list all services
```
kubectl get service
```

list pods
```
kubectl get pod -o wide
```

port forward to specific pod
```
kubectl port-forward $POD_NAME $PORT:$TARGET_PORT
```

shell into container
```
kubectl exec --tty $POD_NAME  -- sh
```

delete everything from the current namespace
```
kubectl delete all --all
```