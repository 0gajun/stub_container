# Stub HTTP server container

Stub HTTP server container for creating ECS services before deploying real apps.

## Usage

Run server with two environment variables, `LISTEN_PORT` and `APP_NAME`.

```
$ go build
$ LISTEN_PORT=3000 APP_NAME=0gajun ./stub_container
```


The server now listens on 3000 port.
Lets' request GET to `localhost:3000`.

```
$ curl http://localhost:3000
stub container for 0gajun service
```

The response includes `APP_NAME` environment variable value.
