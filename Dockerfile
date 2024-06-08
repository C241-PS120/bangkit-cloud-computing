# Stage 1: Build the Go binary
FROM golang:1.22.3 AS build

WORKDIR /app

ARG DB_HOST
ARG DB_NAME
ARG DB_PORT
ARG DB_USER
ARG DB_PASS
ARG ENVIROMENT

ENV DB_HOST ${DB_HOST}
ENV DB_NAME ${DB_NAME}
ENV DB_PORT ${DB_PORT}
ENV DB_USER ${DB_USER}
ENV DB_PASS ${DB_PASS}
ENV ENVIROMENT ${ENVIROMENT}

COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage 2: Create the final image
FROM alpine:latest AS release

WORKDIR /app

COPY --from=build /app/app .
RUN apk --no-cache add ca-certificates tzdata

EXPOSE 8080

ENTRYPOINT ["/app/app"]
