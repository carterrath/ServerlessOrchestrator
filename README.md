# ServerlessOrchestrator

## Accessing PostgreSQL 
1. Go to PostgreSQL App. 
2. Click your personal server. A terminal will appear.
3. Use the folowing command to look at tables in your database: \dt
4. Then to enter into the microservices table use: \d microservices
5. To view the rows of the table use: SELECT * FROM microservices;

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

## React setup
Install node.js and npm
    brew install node
Navigate to program directory
    npm install react-router-dom

## Dependencies
    npm install @mui/material @mui/icons-material @emotion/styled @emotion/react
    npm install --save-dev @types/react @types/react-dom


