package main

import (
	"netcl/configs"
	"netcl/internal/app/cmd"
)

var config = configs.GetRootConfig()

func main() {
	cmd.Execute()
}
