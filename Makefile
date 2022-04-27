all: build

tests:
	go clean -testcache
	go test ./...

build:
	# local build
	CGO_ENABLED=0 go build -ldflags "-s -w" -o build/ip-link-monitor-daemon

	# raspberry pi build
	CGO_ENABLED=0 GOARCH=arm64 go build -ldflags "-s -w" -o build/ip-link-monitor-daemon.arm64

clean:
	unlink build/ip-link-monitor-daemon
	unlink build/ip-link-monitor-daemon.arm64

.PHONY: tests build clean
