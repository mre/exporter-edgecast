# Information on Edgecast and its API

## Platforms
Edgecast maintains several different "platforms" for different use-cases.
They are identified by an integer value and passed to the API url.
Those are the available platforms:
| name  | value | usage/data|
|-------|-------|-----------|
| flash | 2     | unknown   |
| http_large | 3 | bulk objeccts (>= 300KB) stored on an ordinary disk |
| http_small | 8 | smal objects (< 300KB) stored on SSD storage|
| adn | 14 | Application Delivery Network (ADN)

## metric keys that can be queried
- bandwidth
- connections
- cachestatus
- statuscode

## Request
### Default Headers:
- Authorization: <edgecast_token>
- Accept: application/json
- Content-Type: application/json

## Storage
### store to timeseries database using Graphite

