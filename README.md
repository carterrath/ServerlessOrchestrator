# ServerlessOrchestrator

## SQLite Database Steps
1. Download SQLite3 extension on VS Code.
2. In terminal, run "go get github.com/mattn/go-sqlite3"
3. In terminal, run "go get github.com/go-git/go-git/v5"

## Run Minikube
Run minikube tunnel in a separate terminal
minikube start
minikube tunnel

## Run Serverless Orchestrator
go run main.go
enter any microservice name
Paste repository link for microservice: https://github.com/ruthijimenez/service-catalog
- This will add the microservice to the sqlite database
go run main.go
enter same microservice name
enter yes

