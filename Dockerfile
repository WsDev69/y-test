FROM golang:1.14-alpine as build-env
# All these steps will be cached
RUN mkdir /y-test
WORKDIR /y-test
COPY go.mod .
RUN go mod download
COPY . .

# Build the binary
RUN GOOS=linux GOARCH=arm go build -o /go/bin/y-test cmd/test/main.go
FROM scratch
COPY --from=build-env /go/bin/y-test /go/bin/y-test
EXPOSE 80
ENTRYPOINT ["/go/bin/y-test"]