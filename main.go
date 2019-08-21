package main

import (
	"github.com/fwiedmann/heartbeat/cmd"
)

// HeartbeatVersion holds current version
var HeartbeatVersion string

func main() {
	cmd.Execute(HeartbeatVersion)
}
