# UserService RESTful Go server

## Configuration

Create a config.yaml file and use [cfgcrypt][29f7816d] to encrypt it.

Example config file
```
db:
        port: 27017
        server: mongodb
        database: databaseName
        username: databaseUserName
        password: #{{someDatabasePassword}}#
log:
        path: /log/server.log
        level: defaultLogLevel
http:
        hostname: localhost
        port: 80
```

### Example configuration file
```
db:
        server: localhost
        database: application
        username: application
        password: #{{somePassword}}#
log:
        path: log/server.log
        level: info
http:
        hostname: localhost
        port: 8080
```
## Prerequisites
-   [Docker][337bc096]
-   [Docker Compose][8806988c]
-   [Habitus][03481138]


## Build

`habitus`

## Run

`docker-compose up`

## Development

To add source to go path, `cd` to this folder and run `export GOPATH=$GOPATH:`pwd``

[29f7816d]: https://github.com/BernardIgiri/cfgcrypt "Configuration file encryption utility"
[8806988c]: https://docs.docker.com/compose/ "Docker Compose Utility"
[337bc096]: https://docs.docker.com "Docker Paltform"
[03481138]: http://www.habitus.io "Habitus Build Flow Tool For Docker"
