setup:
	go mod download
	go mod vendor
run: local
	./target/local
local:
	go build -mod vendor -o target/local
build: setup
	go build -mod vendor -o target/arc
