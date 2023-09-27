package main

import (
	"dpcli/cfg"
	"dpcli/cmd"
)

func main() {
	cfg.ReadConfig()
	cmd.Execute()
}
