# README #

Migrating a monolithic PHP application to Microservices written in Go

### What is this repository for? ###

* Bachelor Thesis Project
* 1.0.0

### Package Management
* This project uses glide as package manager
* versions are tracked in glide.lock
* glide settings are included in glide.yaml
* get glide here: https://glide.sh/

### make a call to serving microservice:
- via: curl -XPOST -d'{"s":"jfk"}' localhost:8080/
- via postman: 
    + URL: localhost:8080/
    + Body (raw): {"s":"jfk"}
