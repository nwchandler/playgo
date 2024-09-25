package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	os.Exit(Main())
}

// Initializer initializes a new go module so that users can
// quickly get a sandbox environment for their Golang development
// needs.
type Initializer struct {
	Stdin  io.Reader
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

func (i Initializer) Init() error {
	if i.Module == "" {
		name := left[rand.Intn(len(left))] + "_" + right[rand.Intn(len(right))]
		i.Module = name
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

func Main() int {
	init := Initializer{
		Stdin:     os.Stdin,
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

const main_file = `package main

import "fmt"

func main() {
	fmt.Println("Hello, Play-Go!")
}
`

var right = []string{
	"aardvark",
	"alligator",
	"alpaca",
	"ape",
	"anteater",
	"baboon",
	"badger",
	"bald_eagle",
	"barnacle",
	"barracuda",
	"beetle",
	"bird",
	"butterfly",
	"capybara",
	"cardinal",
	"cat",
	"chicken",
	"cow",
	"dog",
	"donkey",
	"duck",
	"goat",
	"goose",
	"guinea_pig",
	"horse",
	"pig",
	"rabbit",
	"sheep",
}

var left = []string{
	"academic",
	"acclaimed",
	"accomplished",
	"acrobatic",
	"active",
	"actual",
	"adept",
	"admirable",
	"adorable",
	"adventurous",
	"affectionate",
	"agile",
	"agreeable",
	"amazing",
	"ambitious",
	"amusing",
	"angelic",
	"animated",
	"apt",
	"artistic",
	"astonishing",
	"athletic",
	"authentic",
	"awesome",
	"beloved",
	"beneficial",
	"better",
	"best",
	"blushing",
	"bold",
	"brave",
	"brilliant",
	"busy",
	"carefree",
	"celebrated",
	"charming",
	"cheerful",
	"cheery",
	"classic",
	"clean",
	"clever",
	"colorful",
	"compassionate",
	"competent",
	"composed",
	"considerate",
	"content",
	"cooperative",
	"courageous",
	"courteous",
	"creative",
	"cuddly",
	"dapper",
	"dazzling",
	"decisive",
	"definitive",
	"dependable",
	"determined",
	"devoted",
	"dutiful",
	"earnest",
	"educated",
	"elaborate",
	"elegant",
	"elementary",
	"enchanting",
	"energetic",
	"enlightened",
	"esteemed",
	"exalted",
	"excellent",
	"exemplary",
	"fabulous",
	"famous",
	"fancy",
	"fantastic",
	"flashy",
	"forthright",
	"friendly",
	"generous",
	"gentle",
	"genuine",
	"graceful",
	"gracious",
	"grandiose",
	"gregarious",
	"handsome",
	"harmonious",
	"helpful",
	"honest",
	"honorable",
	"idealistic",
	"imaginative",
	"impassioned",
	"incomparable",
	"infinite",
	"intelligent",
	"interesting",
	"intrepid",
	"jovial",
	"jubilant",
	"kaleidoscopic",
	"kindhearted",
	"legitimate",
	"lighthearted",
	"likable",
	"lively",
	"magnificent",
	"majestic",
	"memorable",
	"merry",
	"mysterious",
	"noteworthy",
	"novel",
	"optimistic",
	"opulent",
	"outgoing",
	"outstanding",
	"overjoyed",
	"peaceful",
	"perfect",
	"pleasant",
	"polite",
	"positive",
	"prestigious",
	"quintessential",
	"reliable",
	"remarkable",
	"scholarly",
	"scientific",
	"serene",
	"silly",
	"sparkling",
	"spirited",
	"studious",
	"subtle",
	"thankful",
	"thoughtful",
	"thunderous",
	"treasured",
	"tremendous",
	"trustworthy",
	"truthful",
	"unselfish",
	"upbeat",
	"valuable",
	"verifiable",
	"vibrant",
	"victorious",
	"virtuous",
	"warmhearted",
	"whimsical",
	"wise",
	"worthwhile",
	"zany",
}
