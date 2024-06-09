# Stage 1: Build the Go binary
FROM golang:1.22.3 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage 2: Create the final image
FROM alpine:latest AS release

WORKDIR /app

COPY --from=build /app/app .
RUN apk --no-cache add ca-certificates tzdata

EXPOSE 8080

ENTRYPOINT ["/app/app"]
