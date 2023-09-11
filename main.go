package main

import (
	"fmt"

	"github.com/prakashdwasdil/cryptit/decrypt"
	"github.com/prakashdwasdil/cryptit/encrypt"
)

func main() {
	fmt.Println("Encrypted String:- ", encrypt.Nimbus("KodeKcloud"))
	fmt.Println("Decrypted String:- ", decrypt.Nimbus("NrghNforxg"))
}
