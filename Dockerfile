# multi stage https://docs.docker.com/build/building/multi-stage/#use-multi-stage-builds

# build stage
FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY cmd/goserverandom/*.go ./
COPY internal/app/*.go ./internal/app/
COPY internal/web/*.go ./internal/web/

RUN go build -o /docker-goserverandom

# package stage
FROM golang:1.20-alpine

COPY --from=0 /docker-goserverandom /
EXPOSE 8080

ENTRYPOINT /docker-goserverandom 8080