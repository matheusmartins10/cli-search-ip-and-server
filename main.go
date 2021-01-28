package main

import (
	"cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Started")

	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}