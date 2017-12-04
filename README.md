# README #

## Edgecast-Collector

### What is this repository for? ###

* Bachelor Thesis Project
* 2.0.0

### Package Management
* This project uses **dep** as package manager
* versions are tracked in Gopkg.lock
* glide settings are included in Gopkg.toml
* get dep here: [https://github.com/golang/dep](https://github.com/golang/dep)

### test
- GET on [http://localhost:80/metrics](http://localhost:80/metrics)
- via Prometheus from local directory: ```prometheus -config-file=prometheus.yml```

### static analysis
- from local directory: ```make lint``` (uses gometalinter, downloads and installs it in case of absence)

### build
- from local directory: ```make build``` (builds for Windows or Unix, after checking ```$(OS),Windows_NT```)

### run
- from local directory: ```./bin/main``` (Unix) or ```.\bin\main.exe``` (Windows)
- via docker: 
    + build docker image: ```make docker```
    + run docker image: ```(sudo) docker run -P trivago/monitoring:edgecast-v1```
    
### view exposed metrics:
- via Browser on the same machine: visit [http://localhost:80/metrics](http://localhost:80/metrics)
    + via Browser on different machine: change "localhost" to endpoint address
- via existing Prometheus server installation: 
    + start new server locally using the provided configuration file:
        * ```prometheus -config-file=prometheus.yml```
        * view results here: [http://localhost:9090](http://localhost:9090)
    + copy & paste job from provided prometheus.yml to running server's configuration to scrape the service metrics