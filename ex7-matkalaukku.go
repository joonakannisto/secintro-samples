package main

import (
	"fmt"
	"math/big"
)

func main() {

	e := big.NewInt(65537)
	
	n := big.NewInt(1)
	n,success := n.SetString("9539961903522275212673483811017281691831807188661252864554079397202586017541",10)
	if (!success) {
	fmt.Printf("Failed to set the public modulus")
	return
	}
	matkalaukku := big.NewInt(1234)
	matkalaukku = matkalaukku.Exp(matkalaukku,e,n)
	fmt.Printf("The combination is %s \n", matkalaukku.Text(10))
}

