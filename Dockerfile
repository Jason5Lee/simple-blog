FROM golang:1.19.3-alpine3.16 AS build

WORKDIR /build

COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o simple-blog .

FROM alpine:3.16
WORKDIR /
COPY --from=build /build/simple-blog .
EXPOSE 8080
ENTRYPOINT [ "/simple-blog" ]
