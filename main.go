package main

import (
	"github.com/rene00/wtz/cmd"
)

const version = "0.0.5"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
