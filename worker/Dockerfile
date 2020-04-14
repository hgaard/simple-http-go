##################################################################### Builder Stage                                                    # ##################################################################### Start from the golang:alpine image with the latest version of Go installed
FROM golang:alpine AS builder

# Create WORKDIR using project's root directory
WORKDIR /go/src/simple-http-go

# Copy the local package files to the container's workspace
# in the above created WORKDIR
ADD . .

# Build the api-server inside the container
RUN go build -o worker

##################################################################### Final Stage                                                      # ##################################################################### Pull golang alpine image (very small image, with minimum needed to run Go)
FROM alpine

# Create WORKDIR
WORKDIR /app

# Copy app binary from the Builder stage image
COPY --from=builder /go/src/simple-http-go .

# Run the userServer command by default when the container starts.
ENTRYPOINT ./worker
