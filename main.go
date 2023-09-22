package main

import (
	"flag"
	"fmt"
	"minism/cmds"
	"os"
)

func main() {

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("🔴 invalid command")
		os.Exit(0)
	}

	command := flag.Args()[0]

	errCmd := cmds.Parse(command, flag.Args()[1:])
	
	if errCmd != nil {
		fmt.Println(errCmd)
		os.Exit(1)
	}

}
