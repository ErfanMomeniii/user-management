FROM golang:1.20 AS build

ENV GO111MODULE=on \
GOOS=linux \
GOARCH=amd64 \
GOPROXY=$GOPROXY

RUN mkdir -p /src

WORKDIR /src

COPY go.mod go.sum config.yaml /src/
COPY . /src
RUN make build-static-vendor-linux

FROM debian:bullseye AS local

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}"

WORKDIR /app

COPY --from=build /src/user-management /app
COPY --from=build /src/config.yaml /app

CMD ["./user-management", "start"]
