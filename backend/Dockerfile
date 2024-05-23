# Use the official Golang image as the base image
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Install dependencies
RUN apt-get update && apt-get install -y \
    bash \
    default-mysql-client \
    wget \
    && wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && apt-get install -y ./google-chrome-stable_current_amd64.deb \
    && rm -f google-chrome-stable_current_amd64.deb

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go apps
RUN go build -o /app/csvdownloader ./cmd/csvdownloader/main.go
RUN go build -o /app/csvtomysql ./cmd/csvtomysql/main.go
RUN go build -o /app/main ./main.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the the main executable
CMD ["./main"]
