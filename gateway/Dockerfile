FROM golang:1.19-buster as builder

# Create and change to the cert directory.
WORKDIR /cert

# Copy local cert
COPY ./cert/ ./

# Create and change to the app directory.
WORKDIR /app

# Expose the port that the application will listen on
EXPOSE 3000

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY ./go.* ./
RUN go mod download

# Copy local code to the container image.
COPY ./ ./

# Build the binary.
RUN go build -o /server

# Set environment variables
ENV BUILD_ENV=dev

# Run the application
CMD ["/server"]
