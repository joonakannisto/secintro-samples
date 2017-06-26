package main

import (
	"fmt"
	//"encoding/hex"
	"encoding/base64"
)

func main() {
	bStr := "WUVMTE9XIFNVQk1BUklORQ=="
	b,_ := base64.StdEncoding.DecodeString(bStr)
	
	fmt.Printf("%s \n",b)
		
	for index := range b {
		// | binary AND operation 
		b[index] = b[index]|0x20	
	}
	fmt.Printf("%s \n",b)
		
	fmt.Printf("Base64: %s \n", base64.StdEncoding.EncodeToString(b))
}

