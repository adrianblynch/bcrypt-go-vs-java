package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"time"
)

func main() {

	password := []byte("password")

	start := time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 14)
	elapsed := time.Since(start)

	if (err != nil) {
		panic(err)
	}

	fmt.Println("Bcrypt hashing speed check...")
	fmt.Printf("'%s' becomes '%s'\n", string(password), string(hashedPassword))
	fmt.Printf("The hashing took %s\n", elapsed)

}