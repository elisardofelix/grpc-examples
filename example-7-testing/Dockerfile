FROM golang:1.25.1-alpine as build

ENV GONOSUMDB "github.com/cshep4"
ENV GOPRIVATE "github.com/cshep4"

WORKDIR /build

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags='-N -l' -ldflags="-s -w" -o /server ./cmd/server

FROM alpine as release

COPY --from=build /server /server

ENTRYPOINT ["/server"]