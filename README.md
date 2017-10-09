# README #

Migrating a monolithic PHP application to Microservices written in Go

### What is this repository for? ###

* Bachelor Thesis Project
* 1.0.0

### Package Management
* This project uses ~~glide~~ **dep** as package manager
* versions are tracked in ~~glide.lock~~ Gopkg.lock
* glide settings are included in ~~glide.yaml~~ Gopkg.toml
* ~~get glide here: https://glide.sh/~~
* get dep here: https://github.com/golang/dep

### make a call to serving microservice:
- via: curl -XPOST -d'{"s":"jfk"}' localhost:8080/
- via postman: 
    + URL: localhost:8080/
    + Body (raw): {"s":"jfk"}

# the res folder
- contains useful information on files and packages
- contains class/project/Entity-Relationship-Diagrams
    + created with https://github.com/gmarik/go-erd