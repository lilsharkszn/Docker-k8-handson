## Overview
This project demonstrates a complete end-to-end DevOps workflow for containerizing, deploying, and orchestrating a backend application using modern cloud-native tools. 

The system is built around the MuchToDo Backend API, a Go-based service, and showcases how to move from local development to a fully orchestrated Kubernetes environment. 

## This repository highlights real-world DevOps practices including:
- Infrastructure provisioning
- Environment cconfiguration
- Containerization
- Multi-service orchestration
- Kubernetes deployment
- Debugging and system validation 

## What This Project Achieves
- Containerizes a production-ready backend using Docker
- Runs supporting services (MongoDB, Redis) via Docker Compose
- Deploys the full system into Kubernetes (KIND cluster)
- Configures persistent storage for MongoDB
- Exposes services using Kubernetes Service and Ingress
-  Validates application health via API endpoints
- Automates environment setup using Ansible
--- 

## Architecture 

Client 
(curl / browser)
        ↓
NGINX Ingress Controller
        ↓
Kubernetes Service (backend-service)
        ↓
Backend Pod (Golang API)
        ↓
MongoDB Pod (Database) 

--- 

## Tech Stack 

Core Stack
- AWS
- Terraform
- Ansible
- Docker
- Kubernetes (KIND) 

Backend Stack
- Go (Gin)
- MongoDB
- Redis 

Tooling
- Git
- Curl 

NGINX Ingress Controller 

--- 

## Project Structure 

Docker-k8-handson/
.
├── README.md
├── ansible
│   └── setup.yml
├── evidence
│   └── muchToDo-evidence.pdf
├── infra
│   ├── main.tf
│   ├── modules
│   │   ├── ec2
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   └── variables.tf
│   │   ├── security-group
│   │   │   ├── main.tf
│   │   │   ├── output.tf
│   │   │   ├── outputs.tf
│   │   │   └── variables.tf
│   │   └── vpc
│   │       ├── main.tf
│   │       ├── outputs.tf
│   │       └── variables.tf
│   ├── outputs.tf
│   ├── provider.tf
│   └── variables.tf
└── much-to-do
    └── Server
        └── MuchToDo
            ├── Dockerfile
            ├── Makefile
            ├── cmd
            │   └── api
            │       └── main.go
            ├── docker-compose.yaml
            ├── docs
            │   ├── docs.go
            │   ├── swagger.json
            │   └── swagger.yaml
            ├── go.mod
            ├── go.sum
            ├── internal
            │   ├── auth
            │   │   ├── auth.go
            │   │   └── auth_test.go
            │   ├── cache
            │   │   └── cache.go
            │   ├── config
            │   │   └── config.go
            │   ├── database
            │   │   └── database.go
            │   ├── handlers
            │   │   ├── handlers_test.go
            │   │   ├── handlers_test.go.cp
            │   │   ├── health.go
            │   │   ├── todo.go
            │   │   └── user.go
            │   ├── logger
            │   │   └── logger.go
            │   ├── middleware
            │   │   ├── logger.go
            │   │   └── middleware.go
            │   ├── models
            │   │   ├── todo.go
            │   │   └── user.go
            │   └── routes
            │       └── routes.go
            └── kubernetes
                ├── backend
                │   ├── backend configmap.yaml
                │   ├── backend-deployment.yaml
                │   ├── backend-secret.yaml
                │   └── backend-service.yaml
                ├── ingress.yaml
                ├── mongodb
                │   ├── mongodb-deployment.yaml
                │   ├── mongodb-pvc.yaml
                │   ├── mongodb-secret.yaml
                │   └── mongodb-service.yaml
                └── namespace.yaml--- 

## Docker Setup 

### Build Image
docker build -t muchtodo-backend . 

### Run with Docker Compose
docker compose up -d 

### Test API
curl http://localhost:8080/health 

--- 

## Kubernetes Setup (Kind) 

### Create Cluster
kind create cluster --name muchtodo 

### Deploy Application
kubectl apply -f kubernetes/ 

### Verify Resources
kubectl get pods -n muchtodo
kubectl get svc -n muchtodo
kubectl get ingress -n muchtodo 

--- 

## Ingress Access (Important) 

Kind does not expose LoadBalancer externally by default.  
Access is achieved via the Kind node IP and ingress NodePort. 

### Get Node IP
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' muchtodo-control-plane 

Example:
172.19.0.2 

### Get Ingress NodePort
kubectl get svc -n ingress-nginx 

Example:
80:32404/TCP 

### Test API via Ingress
curl -H "Host: muchtodo.local" http://172.19.0.2:32404/health 

Expected Response:
{
  "cache": "disabled",
  "database": "ok"
} 

--- 

## Hosts Configuration 

Add to /etc/hosts: 

127.0.0.1 muchtodo.local 

--- 

## Evidence Folder 

evidence/
├── muchToDo-evidence.pdf

--- 

## Key Learnings 

- Docker multi-stage builds improve production image size
- Docker Compose simplifies local service orchestration
- Kubernetes separates concerns via Pods, Services, and Ingress
- Kind requires Node IP or port-forward for external access
- Ingress enables clean HTTP routing inside Kubernetes 

--- 

## Status 
✔Backend containerized  
✔MongoDB integrated  
✔Docker Compose working  
✔Kubernetes deployment successful  
✔Ingress configured and tested  
✔End-to-end API communication verified  

--- 

## Final Notes 

This project demonstrates a production-like deployment pipeline using local Kubernetes infrastructure, showcasing real-world DevOps engineering practices including containerization, orchestration, and service exposure.
