# Stage 1: Build the Go application
FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Setup a smaller base image using Debian
FROM debian:stable-slim  
WORKDIR /root/

# Install bash, mysql-client, wget, and Google Chrome
RUN apt-get update && apt-get install -y \
    bash \
    default-mysql-client \
    wget \
    && wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && apt install -y ./google-chrome-stable_current_amd64.deb \
    && rm -f google-chrome-stable_current_amd64.deb

# Copy built binary from the build stage
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]
