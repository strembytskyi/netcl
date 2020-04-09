package main

import (
	"netcl/cmd"
	"netcl/configs"
)

var config = configs.GetRootConfig()

func main() {
	cmd.Execute()
}
