# Stage 1: Build the Go application
FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Setup a smaller base image
FROM alpine:latest  
WORKDIR /root/

# Install bash and necessary packages
RUN apk add --no-cache bash mysql-client wget \
    && wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && apk add --no-cache ./google-chrome-stable_current_amd64.deb \
    && rm -f google-chrome-stable_current_amd64.deb

COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]
