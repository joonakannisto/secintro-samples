package main

import (
	"fmt"
	"encoding/hex"
	"crypto/rand"
)

func main() {

	pt:=[]byte("This is no walk in the park")
	rByte:=make([]byte,1)
	_,err := rand.Read(rByte)
	if err != nil { panic(err)}
	for index:=range pt {
		pt[index] = pt[index]^rByte[0]
	}
	fmt.Printf("%s", hex.EncodeToString(pt))
	
}

