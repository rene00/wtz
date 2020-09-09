package main

import (
	"github.com/rene00/wtz/cmd"
)

const version = "0.1.1"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
