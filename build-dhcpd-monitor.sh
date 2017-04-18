echo "Compile and build the application"
docker run --rm -v $(pwd):/usr/src/myapp -w /usr/src/myapp -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 golang:1.8 bash -c "go get -d -v; go build -a --installsuffix cgo -v -o dhcpd-monitor"

echo "Build image"
docker build --no-cache=true -t ictu/dhcpd-monitor .
