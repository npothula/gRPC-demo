# gRPCHelloWorld
This demo HelloWorld Server and Client application is to leverage gRPC (rather than using REST API for communication among the Microservices)

Protocol Buffer has defined for message transmission between gRPC HelloWorld Server and Client.
Refer helloworld/helloworld.proto for more details on the message format.
If you wish to update the messages, complie .proto file with protoc compiler that is shown as below
    protoc --go_out=plugins=grpc:. helloworld/*.proto

Design and Object Modelling:
    Refer design folder for architecture, classDiagram and ProcessFlow of the server and client.

UnitTests:
    Refer unittests in corresponding `_test` files

Build and Deploy:
    Pleass run build_deploy.sh script. It does below tasks.
        1. Builds the Docker image via Docker compose.
        2. Pushes the image into Dcoker Hub(specify your docker-hub registry).
        3. Deploys into Kubernetes Cluster
Check 'is Application Running' commented section in build_deploy.sh for instructions on running the application

TODO:
    Refer TODO file for pending tasks.
