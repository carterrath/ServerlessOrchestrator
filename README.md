# ServerlessOrchestrator 
Install minikube: downloading minikube https://minikube.sigs.k8s.io/docs/start/ #MAKE SURE ITS THE RIGHT ONE
copy n paste curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-arm64 
sudo install minikube-darwin-arm64 /usr/local/bin/minikube -> TERMINAL
run ' go run main.go ' 
Interact with the cluster 'minikube kubectl -- get po -A' 
minikube start 
