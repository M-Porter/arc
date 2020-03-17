setup:
	go mod download
	go mod vendor
run:
	go build -mod vendor -o target/local
	./target/local
build:
	go mod download
	go mod vendor
	go build -mod vendor -o target/arc
