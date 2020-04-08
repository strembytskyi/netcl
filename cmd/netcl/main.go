package main

import (
	"fmt"
	"netcl/configs"
)

var config = configs.GetRootConfig()

func main() {
	fmt.Println(config.Short)
}
