pipeline {
   agent any
   environment {
       REGISTRY = 'npothula/docker-hub'
       registryCredential = 'dockerhub'
       BUILD_NAME = 'grpc-demo-helloworld'
   }
   stages {
       stage('Build') {
           steps {
               sh 'docker-compose -f "docker-compose.yml" build'
           }
       }
       stage('Publish') {
           steps{
               sh 'docker-compose -f "docker-compose.yml" push'
           }
       }
       stage('Deploy') {
           steps{
               sh 'kubectl apply -f deployment.yml'
           }
       }
   }
}