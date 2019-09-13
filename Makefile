all: build run

testVersion = 1.0.0

build:
	CGO_ENABLED=0 go build -o heartbeat -ldflags "-X main.HeartbeatVersion=$(testVersion)"

run: build
	sudo ./heartbeat --loglevel debug

build-docker:
	docker build -t heartbeat .

run-docker: build-docker
	docker run -it -v $$PWD/config.yaml:/code/config.yaml -p 8080:8080 -p 9100:9100 heartbeat