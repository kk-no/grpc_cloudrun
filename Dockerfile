# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder
ENV GO111MODULE=on
RUN apk add --no-cache ca-certificates git

# Copy local code to the container image.
WORKDIR /cloudrun
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o cloudrun

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /cloudrun/cloudrun /cloudrun

# Run the web service on container startup.
CMD ["/cloudrun"]