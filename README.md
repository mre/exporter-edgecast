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
- GET on http://localhost:8080/metrics
- via Prometheus: prometheus -config-file=res/prometheus.yml

# the res folder
- contains useful information on files and packages
- contains class/project/Entity-Relationship-Diagrams
    + created with [go-erd](https://github.com/gmarik/go-erd, "https://github.com/gmarik/go-erd")
    + ```go-erd -path . | dot -Tsvg > out.svg```
- call graph created with [go-callvis](https://github.com/TrueFurby/go-callvis,"https://github.com/TrueFurby/go-callvis")
    + ```go-callvis [OPTIONS] <main pkg> | dot -Tpng -o output.png```
    + grouped by package using option ```-group pkg```