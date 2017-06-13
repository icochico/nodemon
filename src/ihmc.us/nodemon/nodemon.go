package main

import (
	"ihmc.us/nodemon/cmd"
	"github.com/CrowdSurge/banner"
	"fmt"
)

//variables for versioning
var (
	Version   string
	BuildTime string
	GitHash   string
)

func main() {
	fmt.Println(banner.PrintS("nodemon"))
	fmt.Println(cmd.DescrSignature)
	fmt.Println("Version: " + Version)
	fmt.Println("Build Time: " + BuildTime)
	fmt.Println("Git Hash: " + GitHash)

	//run NodeMon by executing Cobra root command
	cmd.Execute()
}
