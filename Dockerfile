# Build stage of gRPCHelloWorld at intermediate container
FROM golang:1.14.1-alpine3.11 AS dev_img

ARG HOME_PATH="/home/gRPCHelloWorld"
ARG GO_APP_PATH="/go/src/gRPCHelloWorld"
ARG APP_BIN_PATH="/home/gRPCHelloWorld/bin"

RUN apk update && apk upgrade \
    && apk add --no-cache ca-certificates

# copying the source code and build in the container
RUN mkdir -p $GO_APP_PATH
WORKDIR $GO_APP_PATH
COPY . $GO_APP_PATH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -o $APP_BIN_PATH/svcGRPCHelloWorld $GO_APP_PATH/src/server
RUN chmod +x $APP_BIN_PATH/svcGRPCHelloWorld
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -o $APP_BIN_PATH/grpcHelloWorldClient $GO_APP_PATH/src/client
RUN chmod +x $APP_BIN_PATH/grpcHelloWorldClient


# final stage
FROM alpine:3.11 AS prod_img

RUN mkdir -p /home/gRPCHelloWorld \
    && apk --no-cache add ca-certificates

WORKDIR /home/gRPCHelloWorld
COPY --from=dev_img /home/gRPCHelloWorld .

CMD ["/home/gRPCHelloWorld/bin/svcGRPCHelloWorld"]