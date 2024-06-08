# Stage 1: Build the Go binary
FROM golang:1.22.3 AS build

WORKDIR /go/src/github.com/C241-PS120/bangkit-cloud-computing

ARG DB_HOST
ARG DB_NAME
ARG DB_PORT
ARG DB_USER
ARG DB_PASS

ENV DB_HOST ${DB_HOST}
ENV DB_NAME ${DB_NAME}
ENV DB_PORT ${DB_PORT}
ENV DB_USER ${DB_USER}
ENV DB_PASS ${DB_PASS}

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

# Stage 2: Create the final image
FROM alpine:latest AS release
WORKDIR /app
COPY --from=build /go/src/github.com/C241-PS120/bangkit-cloud-computing/main .

RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/main

EXPOSE 8080

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/app/main"]
