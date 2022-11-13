FROM golang:1.19.3-alpine3.16

COPY ./ /test

WORKDIR /test

RUN go mod download
ENTRYPOINT [ "go", "test", "-tags=integration", "./infra/integration_test", "-v", "-count=1" ]
