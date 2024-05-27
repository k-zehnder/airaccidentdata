# Use the official Debian-based Go image
FROM golang:1.20

# Install necessary packages
# - default-mysql-client: Required for MySQL database interactions
# - wget: Used to download files, including Chrome
RUN apt-get update && apt-get install -y \
    default-mysql-client \
    wget

# Download and install Google Chrome
# Chrome is needed for chromedp, which automates browser interactions to fetch the Federal Aviation Administration CSV download link.
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && apt install -y ./google-chrome-stable_current_amd64.deb \
    && rm -f google-chrome-stable_current_amd64.deb

# Install swag CLI for generating Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Set the working directory for the application
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Generate Swagger docs
RUN swag init

# Verify Swagger docs generation
RUN test -f docs/swagger.json && test -f docs/swagger.yaml

# Ensure init-db.sh and run-csv-processors.sh are executable
RUN chmod +x /app/bin/init-db.sh /app/bin/run-csv-processors.sh

# Expose port 8080
EXPOSE 8080

# Run the initialization script and start the Go application
CMD ["/app/bin/init-db.sh"]
