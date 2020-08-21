FROM golang:1.11 as build

ENV GO111MODULE on

WORKDIR /go/src/app
COPY . .

RUN go install -a -tags netgo -installsuffix netgo -ldflags='-extldflags="static"' -v .

FROM alpine
RUN apk --no-cache add --update curl

# run as non-root user for kubernetes (especially pod security policy)
# nobody:nobody
USER 65534:65534
COPY --chown=nobody:nobody --from=build /go/bin/stub_container /
COPY --chown=nobody:nobody server.crt server.key /

CMD ["/stub_container"]
