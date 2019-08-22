all: build run

testVersion = 1.0.0

build:
	CGO_ENABLED=0 go build -o heartbeat -ldflags "-X main.HeartbeatVersion=$(testVersion)"

run:
	sudo ./heartbeat --loglevel debug

build-docker:
	docker build -t heartbeat .