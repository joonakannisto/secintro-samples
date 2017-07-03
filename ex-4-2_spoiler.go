package main

import (
	"fmt"
	"encoding/hex"
	"crypto/rand"
)

func main() {

	b := []byte("This is some secret text")
	
	// Password is some random byte 
	pass := make([]byte,len(b))
	_,err := rand.Read(pass)
	if (err != nil) {panic(err)}
		
	for index := range b {
		// ^ denotes bytewise XOR operation 
		b[index] = b[index]^pass[index]	
	}
	
	
	// Output should look like garbage
	fmt.Printf("%s \n",b)
	
	fmt.Printf("%s \n", hex.EncodeToString(b))
	
	
	
	freq := []byte("ETAOIN SHRDLU")
	
	best := 0
	solution :=0 
	// go through all possible keys
	for i:=0; i<256; i++ {
		current :=0
		// lower case comparison example
		if (b[5] ^ byte(i)|0x20 == freq[2]|0x20 ) {
			fmt.Printf("A")
			// Vokaali A ostettu, 10 pistettä rohkelikolle 
			current +=10
		}
		if (current > best) {
			best = current 
			solution = i
		}
	}
	fmt.Printf("%d", solution)
}

