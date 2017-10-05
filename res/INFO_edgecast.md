# Information on Edgecast and its API

## Platforms
Edgecast maintains several different "platforms" for different use-cases.
They are identified by an integer value and passed to the API url.
Those are the available platforms:
| name       | value | usage/data                                          |
|------------|-------|-----------------------------------------------------|
| flash      | 2     | unknown                                             |
| http_large | 3     | bulk objeccts (>= 300KB) stored on an ordinary disk |
| http_small | 8     | small objects (< 300KB) stored on SSD storage       |
| adn        | 14    | Application Delivery Network (ADN)                  |

## metric keys that can be queried (with example values)
- bandwidth
    + ```{"Result":42.42}```
- connections
    + ```{"Result":1234.1234}```
- cachestatus
    + ```[{"CacheStatus": "<StatusName>", "Connections": <int value>}, ...]```
        * StatusName = TCP_HIT, TCP_EXPIRED_HIT, TCP_MISS, ...
        * see example fixtures [here](./fixtures/cachestatus.json) 
- statuscode
    + ```[{"Connections": <int-value>, "StatusCode": "<http-statuscode>"}, ...]```
        * see example fixtures [here](./fixtures/statuscode.json) 

## Request
### Default Headers:
- Authorization: <edgecast_token>
- Accept: application/json
- Content-Type: application/json

## Storage
### store to timeseries database using Graphite

