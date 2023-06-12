# Set the default value for the GOARCH build argument.
ARG GOARCH=amd64

############################
# STEP 1 prepare the source
############################
FROM golang:1.20-alpine AS builder

# Set the environment variables for the go command:
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=$GOARCH

# Set the Asia/Bangkok timezone
RUN apk --no-cache add tzdata
ENV TZ=Asia/Bangkok

# Create a non-root user and group
RUN adduser -D -g '' appuser

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Copy go.mod, go.sum and dependencies if exists
COPY go.mod go.sum vendor? /src/

# Check if the "vendor" folder exists on the host
RUN test -d vendor || go mod tidy

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

# Import the timezone data
COPY --from=builder /usr/share/zoneinfo/Asia/Bangkok /etc/localtime

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the compiled Go application and set ownership to appuser
COPY --from=builder --chown=appuser:appuser /bin/app /bin/app

# Switch to the non-root user
USER appuser

# Start the application
CMD ["/bin/app"]

# docker compose example
# version: '3'
# services:
#   app:
#     build:
#       context: .
#       dockerfile: Dockerfile
#     volumes:
#       - /dev/ttyACM0:/dev/ttyACM0
#     devices:
#       - /dev/ttyACM0