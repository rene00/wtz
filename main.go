package main

import (
	"github.com/rene00/wtz/cmd"
)

const version = "0.1.2"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
