FROM golang:1.13 AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o picsorter

FROM alpine:3.10

WORKDIR /app
COPY --from=build /go/src/app/picsorter /app/picsorter

ENTRYPOINT ["./picsorter"]
