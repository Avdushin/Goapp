package utils

import (
	"errors"
	"fmt"
	"goapp/pkg/consts"
	"log"
	"os"
	"runtime"

	cp "github.com/otiai10/copy"
)

func CreateApp(name string) {
	fmt.Printf("Creating %s%s%s app...", consts.Bold, name, consts.Reset)
	// Create th app folder
	CreateAppFolder(name)
	// Create File structure
	CreateStructure(name)
	// Create General files
	CreateTemplFie(fmt.Sprintf("%s/cmd/main.go", name), fmt.Sprintf("%s/cmd/%s.go", name, name))
	fmt.Printf("\n\ncd %s%s%s\ngo mod init\n", consts.Cyan, name, consts.Reset)
}

func CreateAppFolder(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if e := os.Mkdir(path, os.ModePerm); e != nil {
			log.Fatal(e)
		}
	}
}

func CreateStructure(name string) {
	if runtime.GOOS == "darwin" {
		if err := cp.Copy("/usr/local/bin/Template", name); err != nil {
			log.Fatal("не работайт")
		}
	} else if runtime.GOOS == "linux" {
		if err := cp.Copy("/usr/bin/Template", name); err != nil {
			log.Fatal("не работайт")
		}
	}

}

func CreateTemplFie(file, name string) {
	if e := os.Rename(file, name); e != nil {
		log.Fatal(e)
	}
}

// func GoMod(name string) {
// 	cmd := exec.Command("sh", name+"/gomod.sh")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stdin = os.Stdin
// 	cmd.Stderr = os.Stderr
// 	cmd.Run()
// }
