# Builds the Docker image via Docker compose, push into Dcoker Hub and then Deploy into Kubernetes Cluster

export REGISTRY=npothula/docker-hub
export BUILD_NUMBER=2
export BUILD_NAME=grpc-demo-helloworld

#Build stage
docker-compose -f "docker-compose.yml" build
sleep 1m

#Publish stage
docker push npothula/docker-hub:grpc-demo-helloworld-2
sleep 1m

#Deploy stage
kubectl apply -f deployment.yml

# is Applicaiton running ?
# kubectl get pods
# kubectl exec -it grpc-demo-helloworld-<id> sh
# ./bin/grpcHelloWorldClient
#    output: Greeting: Hello world
# ./bin/grpcHelloWorldClient <msg>
#    output: Greeting: Hello msg

# cleanup stage
# kubectl delete svc,deploy grpc-demo-helloworld