FROM golang:1.11 as build

ENV GO111MODULE on

WORKDIR /go/src/app
COPY . .

RUN go install -a -tags netgo -installsuffix netgo -ldflags='-extldflags="static"' -v .

FROM alpine
COPY --from=build /go/bin/stub_container /
COPY server.crt server.key /
RUN apk --no-cache add --update curl
CMD ["/stub_container"]
