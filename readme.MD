# NOTES

```sh
# create image by name api-server
docker build -t math2do/api-server .

docker push math2do/api-server:latest

# run the image in a container, issue with config
docker run -e ENV=LOCAL -p 8080:8080 math2do/api-server

# for PROD
docker run -e ENV=PROD -p 8080:8080 math2do/api-server

# deploy to k8s cluster
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```
