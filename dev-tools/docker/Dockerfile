FROM golang:1.19-bullseye as build

WORKDIR /go/src/app
ADD . /go/src/app/

RUN go get -d -v ./...
RUN go build -o /go/bin/app cmd/hello.go

###

FROM gcr.io/distroless/base-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]