#download dependencies
go mod vendor

# build amd64
docker build --platform linux/amd64 -f dockerfile -t vm-link2500:0.0.0 .

#build arm64
docker build --platform linux/arm64 --build-arg GOARCH=arm64 -f dockerfile -t vm-link2500:0.0.0-arm64 .

# test
docker run -it vm-link2500:0.0.0