package main

import (
	"fmt"
	"os"

	"actify_srv/internal/db"
	"actify_srv/internal/sys"
)


func main() {
	fmt.Println("Actify Main Server start..")

	args := sys.NewActifyArgs()
	provided := os.Args[1:]
	err := args.Parse(provided)
	if err != nil {
		fmt.Printf("Error: Argument Parsing error.(err: %v)\n", err)
		os.Exit(1)
	}

	fmt.Println("Info: Parse success!!")
	fmt.Println(args)

	database := db.NewPostgresDb()
	err = database.InitializePostgres(args.PostgresConnStr())

	database.Destroy()
	fmt.Println("Info: Actify Main Server finish..")
}