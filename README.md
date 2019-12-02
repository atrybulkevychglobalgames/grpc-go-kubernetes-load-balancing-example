# Golang gRPC services and kubernetes load balancing example

### start with [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/):

```bash
# git clone 
# cd /kubernetes/
# minikube start
# kubectl apply -f ./server-deployment.yml
# kubectl apply -f ./server-service.yml 
# kubectl apply -f ./client-deployment.yml
# k logs -f `k get po --no-headers -o custom-columns=":metadata.name" | grep client`
```
Output example
```bash
...
3 server-556fb75854-mjhdc
4 server-556fb75854-848j7
5 server-556fb75854-gqxq2
6 server-556fb75854-mjhdc
7 server-556fb75854-848j7
8 server-556fb75854-gqxq2
9 server-556fb75854-mjhdc
...
```
Where each line shows which pod gives response ...