all: build run

testVersion = 1.0.0

build:
	go build -o heartbeat -ldflags "-X main.HeartbeatVersion=$(testVersion)"

run:
	sudo ./heartbeat --loglevel debug