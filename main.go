package main

import (
	"distributed-democracy/core"
	"os"
)

func main() {
	nodeRole := os.Args[1]
	switch nodeRole {
	case "master":
		core.GetMasterNode().Start()
	case "worker":
		core.GetWorkerNode().Start()
	default:
		panic("invalid node role")
	}
}
