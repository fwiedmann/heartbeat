all: build run

build:
	go build -o heartbeat main.go

run:
	HEARTBEAT_METRICS_PASSWORD="secret" ./heartbeat --loglevel debug