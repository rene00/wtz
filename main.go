package main

import (
	"wtz/cmd"
)

const version = "0.0.1"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
