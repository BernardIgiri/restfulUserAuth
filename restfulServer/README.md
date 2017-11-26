# RESTful Go server

## Build

```
go get godep
godep save
godep go build
```

## Configuration

Create a config.yaml file and use <https://github.com/BernardIgiri/cfgcrypt> to encrypt it.

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

## Run

`./restfulUserAuth path/to/config.yaml`
