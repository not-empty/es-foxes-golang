package main

import (
	"es-foxes/structs"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Env file not found")
	}

	structs.Variables.FillEnvs()

	ShowFox("symbol.txt")
	os.Mkdir("backup", 0755)

	if len(os.Args) == 1 {
		fmt.Println("Please pass some command, Alloweds:")

		ShowFuns()
		return
	}

	if _, ok := Funs[os.Args[1]]; ok {
		Funs[os.Args[1]](os.Args...)
	} else {
		fmt.Println("Command not found, Alloweds:")

		ShowFuns()
	}

	ShowFox("symbolend.txt")
}
