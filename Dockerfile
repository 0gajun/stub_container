FROM golang:1.11 as build

ENV GO111MODULE on

WORKDIR /go/src/app
COPY . .

RUN go install -v .

FROM gcr.io/distroless/base
COPY --from=build /go/bin/stub_container /
CMD ["/stub_container"]
