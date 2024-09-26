package playgo

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func Main() int {
	init := Initializer{
		Stdout:    os.Stdout,
		Stderr:    os.Stderr,
		Directory: "scratches",
	}
	err := init.Init()
	if err != nil {
		return 1
	}

	return 0
}

// Initializer initializes a new go module so that users can
// quickly get a sandbox environment for their Golang development
// needs.
type Initializer struct {
	Stdout io.Writer
	Stderr io.Writer

	// Directory is the parent directory in which to put
	// the new module. Note that a directory with the name
	// of the module will be created within this directory -
	// it is not the name of the directory where the go.mod
	// will be added.
	Directory string

	// Module is the name of the module that you wish to create.
	// If none is provided, one will be made up.
	Module string
}

// Init initializes a new go module using the configuration of
// the associated [Initializer].
func (i Initializer) Init() error {
	if i.Module == "" {
		i.Module = GenerateRandomName()
	}

	cmd := exec.Command("go", "mod", "init", i.Module)
	path := filepath.Join(i.Directory, i.Module)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Fprintln(i.Stderr, err.Error())
		return err
	}
	cmd.Dir = path

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(i.Stderr, string(out))
		return err
	}

	fmt.Fprintf(i.Stdout, "Created module %q in directory %q\n", i.Module, path)

	mainContent := []byte(main_file)
	mainPath := filepath.Join(path, "main.go")
	err = os.WriteFile(mainPath, mainContent, 0644)
	if err != nil {
		i.Stderr.Write([]byte(out))
		return err
	}

	fmt.Fprintf(i.Stdout, "Created main.go in directory %q\n", path)
	return nil
}

const main_file = `package main

import "fmt"

func main() {
	fmt.Println("Hello, Play-Go!")
}
`
