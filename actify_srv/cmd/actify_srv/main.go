package main

import (
	"fmt"
	"os"

	"actify_srv/internal/sys"
)


func main() {
	fmt.Println(("Actify Main Server start.."))

	args := sys.NewActifyArgs()
	provided := os.Args
	err := args.Parse(provided)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Parse success")
	fmt.Println(args)
}