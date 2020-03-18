setup:
	go mod download
	go mod vendor
clean:
	-rm -rf ./target
run: local
	./target/local
local:
	go build -mod vendor -o target/local
build: setup clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod vendor -o target/arc.darwin-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -o target/arc.linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -mod vendor -o target/arc.linux-arm64
