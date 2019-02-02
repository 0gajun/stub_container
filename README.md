# Stub Container

Stub Container for creating ECS services before deploying real apps.

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

## Feature
* Reachable Test

### Reachable Test

Stub Container can check TCP reachability to any hosts and ports.

If you wanna check whether the container can establish TCP connection to github.com:443, you can use the endpoint `/nettest/<host>/port` like following.

```
# curl http://localhost:3000/nettest/github.com/443
Successfully established tcp connection to github.com:443
```
