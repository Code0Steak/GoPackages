package main

import (
	"fmt"

	"github.com/Code0Steak/GoPackages/decrypt"
	"github.com/Code0Steak/GoPackages/encrypt"
)

func main() {

	eMessage := encrypt.Enc("Mocha Pom!")
	dMessage := decrypt.Dec(eMessage)
	fmt.Print(eMessage, dMessage)
}
