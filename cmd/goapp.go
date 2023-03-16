package main

import (
	"flag"
	"fmt"
	"goapp/pkg/consts"
	"goapp/pkg/utils"
	"os"
)

func main() {
	var h, v bool
	flag.BoolVar(&h, "h", false, "[-h] usage information")
	flag.BoolVar(&v, "v", false, "[-v] show application version")
	flag.Parse()

	values := flag.Args()

	if v {
		fmt.Printf("GOAPP version: %s%s%s\n\n", consts.Yellow, consts.Version, consts.Reset)
	}

	if len(values) == 0 {
		utils.Logo()
		flag.PrintDefaults()
		fmt.Printf("\n\n%sUse %sgoapp %sMyApp%s\n\n", consts.Yellow, consts.Reset, consts.Gray, consts.Reset)
		os.Exit(0)
	} else if len(values) > 1 {
		fmt.Println("Please enter Appkication's name")
	}

	for _, name := range values {

		if len(values) <= 1 && name != "v" && name != "version" {
			utils.Logo()
			utils.CreateApp(name)
			fmt.Printf("\n%sThe %s app has been created!%s\n", consts.Cursive, name, consts.Reset)
		} else if name == "v" || name == "version" || name == "--v" {
			fmt.Printf("GOAPP version: %s%s%s\n\n", consts.Yellow, consts.Version, consts.Reset)
		}
	}
}
