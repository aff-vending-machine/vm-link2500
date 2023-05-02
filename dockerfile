# Set the default value for the GOARCH build argument.
ARG GOARCH=amd64

############################
# STEP 1 prepare the source
############################
FROM golang:1.20-alpine AS builder

# Set the environment variables for the go command:
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=$GOARCH

# Create a non-root user and group
RUN adduser -D -g '' appuser

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Copy go.mod and go.sum and download dependencies (comment below if the project has vendor folder)
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

# switch to the root user to modify device permissions
USER root

# Give appuser permission to access ttyACM0
RUN if [ -e /dev/ttyACM0 ]; then \
        chmod a+rw /dev/ttyACM0 && \
        chown :appuser /dev/ttyACM0; \
    fi

# Switch to the non-root user
USER appuser

# Start the application
CMD ["/bin/app"]