package main

import (
	"github.com/rene00/wtz/cmd"
)

const version = "0.0.4"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
