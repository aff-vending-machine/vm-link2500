############################
# STEP 1 prepare the source
############################
FROM golang:1.20 AS builder

# Set the environment variables for the go command:
# * CGO_ENABLED=1 to build a non-statically-linked executable
# * GO111MODULE=on Force the go compiler to use modules
# * GOOS=linux to run on linuxos
# * GOARCH=arm64 to run on arm64 architecture
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
#
# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# RUN mkdir /user \
#   && echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd \
#   && echo 'nobody:x:65534:' > /user/group

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

#This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line)
# And compile the project
COPY go.mod go.sum /src/

RUN go mod download \
  && go mod verify

# Import the code from the context.
COPY . /src/

#compile the project
RUN go build \
  # -mod=mod \
  # -a \
  # -installsuffix 'static' \
  -o /bin/app \
  /src/cmd/app

# ############################
# # STEP 2 the running container
# ############################
FROM scratch AS final
LABEL maintainer="Tanawat Hongthai <ztrixack.th@gmail.com>"

# Import the user and group files from the first stage.
# COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable.
COPY --from=builder /bin/app /app

# Perform any further action as an unprivileged user.
# USER nobody:nobody

# Run the binary.
ENTRYPOINT ["/app"]
