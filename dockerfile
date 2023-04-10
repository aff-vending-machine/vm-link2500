############################
# STEP 1 prepare the source
############################
FROM golang:1.20-alpine AS builder

# Set the environment variables for the go command:
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64

# Create a non-root user and group
RUN adduser -D -g '' appuser

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum /src/
RUN go mod tidy

# Import the code from the context.
COPY . /src/

# Build the Go application
RUN go build -o /bin/app /src/cmd/app

############################
# STEP 2 the running container
############################
FROM alpine AS runner
LABEL maintainer="Tanawat Hongthai <ztrixack.th@gmail.com>"

# Import the user and group from the builder stage.
COPY --from=builder /etc/passwd /etc/group /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the compiled Go application and set ownership to appuser
COPY --from=builder --chown=appuser:appuser /bin/app /bin/app

# Give appuser permission to access ttyACM0
RUN chown :appuser /dev/ttyACM0 && chmod g+rw /dev/ttyACM0

# Switch to the non-root user
USER appuser

# Start the application
CMD ["/bin/app"]