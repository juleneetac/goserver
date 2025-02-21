### Start minikube
1. Open docker desktop
2. insert command in WSL: minikube start

### Build stage
docker build -t go-server .

docker tag go-server julendk/go-server:1.0.2

docker push julendk/go-server:1.0.2


1. Una vez se tiene el pod y el service corriendo, al ser nodePort se debe hacer forwarding con:
minikube service <nombre service>