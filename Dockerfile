# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/localleon/pingpong

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY ./pingpong/ .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 9111 to the outside world
EXPOSE 9111

# Run the executable
CMD ["pingpong","--config=config-example.yaml"]


