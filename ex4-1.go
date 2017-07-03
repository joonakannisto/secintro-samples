package main

import (
	"fmt"
	"encoding/hex"
	"crypto/sha256"
	"bytes" 
)

func main() {
	plaintext := []byte("Attack at dawn")
	plaintext2 := []byte("Attack at dusk")
	
	ciphertextHex := "1229c8a858f954fe98e1154d71c4"
	
	ciphertext, err := hex.DecodeString( ciphertextHex)
	if (err!=nil) {panic(err)} 
	
	// Write the steps to modify the ciphertext
	// to decrypt to plaintext2 instead of plaintext
	ciphertext2 := ciphertext 

	// Do something to the ciphertext here 

	// You can use the function checkResult to check 
	// your solution
	if (checkResult(ciphertext2)) {
		fmt.Printf("Congratulations, your solution %s was correct, now some generals will %s", hex.EncodeToString(ciphertext2), plaintext2)
	} else {
		fmt.Printf("Your solution is not correct, generals will proceed with %s", plaintext) 
	}
}


func checkResult(ciphertext2 []byte) bool {
	// Tästä käytetään joskus nimeä bittiin sitoutuminen, bit commitment
	// huom. tarkoitus ei ole bruteforcettaa tarkistinta vasten 
	verifySum,err := hex.DecodeString("f2bd6ddb1cd6b78c6bde148febf51d40703a9ba07e4223db90e92eb594d312de")
	if (err!=nil) {panic(err)} 
	sum := sha256.Sum256(ciphertext2) 
	if (bytes.Equal(verifySum[:],sum[:])) {
		return true
	} else {
		return false
	}
		
}

