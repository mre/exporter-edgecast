## Collector-Edgecast

### What is this repository for?

This is a Prometheus Collector for Edgecast CDN.

Collector-Edgecast uses the [edgecast-client](https://github.com/mre/edgecast) created by [Matthias Endler](https://github.com/mre) to fetch metrics from the
EdgeCast CDN API and then transforms and exposes them to be scraped and displayed by [Prometheus](https://prometheus.io/).

### Package Management
* This project uses **dep** as package manager
* versions are tracked in Gopkg.lock
* glide settings are included in Gopkg.toml
* get dep here: [https://github.com/golang/dep](https://github.com/golang/dep)

### Static Analysis
- from local directory: ```make lint``` (uses gometalinter, downloads and installs it in case of absence)

### Build
- from app directory: ```make build``` (builds for Windows or Unix, after checking ```$(OS),Windows_NT```)

### Configure
- You need to set two environment-variables to configure your Edgecast-Account:
    + EDGECAST_ACCOUNT_ID
    + EDGECAST_TOKEN
- e.g. on Linux: ```export EDGECAST_TOKEN=B12AC```

### Run
- from app directory: ```./bin/main``` (Unix) or ```.\bin\main.exe``` (Windows)
- via docker:
    + build docker image: ```make docker```
    + run docker image: ```(sudo) docker run -P trivago/monitoring:edgecast-v1```

### View Exposed Metrics:
- via Browser on the same machine: visit [http://localhost:80/metrics](http://localhost:80/metrics)
    + via Browser on different machine: change "localhost" to endpoint address
- via existing Prometheus server installation: 
    + start new server locally using the provided configuration file:
        * ```prometheus -config-file=prometheus.yml```
        * view results here: [http://localhost:9090](http://localhost:9090)
    + copy & paste job from provided prometheus.yml to running server's configuration to scrape the service metrics