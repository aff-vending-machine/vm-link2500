# Specifies a parent image
FROM golang:1.20-alpine

# Set the environment variables for the go command:
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=arm64

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum /src/
RUN go mod tidy

# Import the code from the context.
COPY . /src/

# Build the Go application
RUN go build -o /bin/app /src/cmd/app

# Specifies the executable command that runs when the container starts
CMD [ "/bin/app" ]