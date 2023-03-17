package utils

import (
	"bytes"
	"errors"
	"fmt"
	"goapp/pkg/consts"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	cp "github.com/otiai10/copy"
)

// Create an Application...
func CreateApp(name string) {
	fmt.Printf("Creating %s%s%s app...", consts.Bold, name, consts.Reset)
	// Create th app folder
	CreateAppFolder(name)
	// Create File structure
	CreateStructure(name)
	// Create Internal folder
	MakeInternal(name, "app")
	MakeInternal(name, "pkg")
	// Create General files
	CreateTemplFie(fmt.Sprintf("%s/cmd/main.go", name), fmt.Sprintf("%s/cmd/%s.go", name, name))
	// MakeFile
	SetMakeFile(name)
	// Done output message
	fmt.Printf("\n\ncd %s%s%s\ngo mod init github.com/%s\n", consts.Cyan, name, consts.Reset, name)
}

// Create App Folder ...
func CreateAppFolder(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if e := os.Mkdir(path, os.ModePerm); e != nil {
			log.Fatal(e)
		}
	}
}

// Create File Structure ...
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

// Rename main.go file ...
func CreateTemplFie(file, name string) {
	if e := os.Rename(file, name); e != nil {
		log.Fatal(e)
	}
}

// Create internal directory
func MakeInternal(name, path string) {
	err := os.MkdirAll(fmt.Sprintf("%s/internal/%s", name, path), 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

// Makefile variables
func SetMakeFile(name string) {
	input, err := ioutil.ReadFile(fmt.Sprintf("%s/Makefile.txt", name))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte("YOUR_APP_NAME"), []byte(name), -1)

	if err = ioutil.WriteFile(fmt.Sprintf("%s/Makefile", name), output, 0777); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func GoMod(name string) {
// 	cmd := exec.Command("sh", name+"/gomod.sh")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stdin = os.Stdin
// 	cmd.Stderr = os.Stderr
// 	cmd.Run()
// }
