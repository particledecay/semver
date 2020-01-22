# multi-stage dockerfile
FROM golang:alpine AS builder

RUN apk update && \
    apk add --no-cache git

RUN mkdir /app

WORKDIR /app
COPY . /app

RUN go mod vendor && \
    GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /usr/local/bin/semver

# final image
FROM scratch

COPY --from=builder /usr/local/bin/semver /usr/local/bin/semver

ENTRYPOINT [ "/usr/local/bin/semver" ]