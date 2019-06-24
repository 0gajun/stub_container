# Stub Container

Stub Container for creating ECS services before deploying real apps.

## Usage

Run server with two environment variables, `LISTEN_PORT` and `APP_NAME`.

```
$ docker build -t stub_container .
$ docker run -e LISTEN_PORT=3000 -e APP_NAME=0gajun -p 3000:3000 stub_container
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
* HTTPS support

### Reachable Test

Stub Container can check TCP reachability to any hosts and ports.

If you wanna check whether the container can establish TCP connection to github.com:443, you can use the endpoint `/nettest/<host>/port` like following.

```
# curl http://localhost:3000/nettest/github.com/443
Successfully established tcp connection to github.com:443
```

### HTTPS support
Stub Container can listen on HTTPS using a self-signed certificate.
If `LISTEN_HTTPS` environment variable is set to `true` string value, Stub container uses HTTPS instead of HTTP.

```
$ docker run -e LISTEN_HTTPS=true -e LISTEN_PORT=3000 -e APP_NAME=0gajun -p 3000:3000 stub_container
$ curl -k https://localhost:3000
stub container for 0gajun service
```
