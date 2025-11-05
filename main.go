package main

import (
	"github.com/mrmahile/gitxpose/cmd"
	"github.com/mrmahile/gitxpose/setup" // Import the init package
)

func main() {
	setup.InitConfig() // Call the config setup function before running the commands
	cmd.Execute()
}
