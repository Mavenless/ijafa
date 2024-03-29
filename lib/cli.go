package lib

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Update() {
	PrintInfo("Updating...")
	_, err := exec.Command("go", "get", "-u", "github.com/mavenless/ijafa").Output()

	if err != nil {
		log.Fatal(err)
	}

	PrintDone("Download finished !")
	os.Exit(0)
}

func Help() {
	fmt.Println("Help:")
	fmt.Println("ijafa -h")
	fmt.Println("ijafa -u")
	fmt.Println("ijafa -d <directory>")
	fmt.Println("ijafa <file>")
}

/**
 *	Cli
 *	Handle the command line interface
 */
func Cli() {
	update := flag.Bool("u", false, "update ijafa")
	help := flag.Bool("h", false, "help")
	directory := flag.String("d", "", "directory")

	flag.Parse()

	if *update {
		Update()
	} else if *help {
		Help()
	} else if *directory != "" {
		ConvertFileDirectory(*directory)
	} else {
		if len(os.Args) == 1 {
			fmt.Println("No file specified")
			fmt.Println("Use -h for help")
		} else {
			for i := 1; i < len(os.Args); i++ {
				PrintInfo("Converting " + os.Args[i] + "...")
				ConvertFile(os.Args[i])
			}
			// ConvertFile(os.Args[1])
		}
	}
}
