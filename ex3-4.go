package main

import (
	"fmt"
	"github.com/nbutton23/zxcvbn-go"
)

func main() {
	
	fmt.Println(zxcvbn.PasswordStrength("abcd",[]string{}))
}

