package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, 14)
}

func CheckPasswordHash(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

func Example() {
	//"Correct" Password
	myPassword := []byte("Password9001")

	//"Incorrect" Password
	myPassw0rd := []byte("Password9000")

	//Hash the correct password
	hashed, err := HashPassword(myPassword)
	if err != nil {
		return
	}

	//Print the hash it generated
	fmt.Println(string(hashed))

	//Checking hash against wrong password
	checked := CheckPasswordHash(hashed, myPassw0rd)
	if checked != nil {
		fmt.Println(checked)
	} else {
		fmt.Println("Matched successfully")
	}

	//Checking hash against correct password
	checked = CheckPasswordHash(hashed, myPassword)
	if checked != nil {
		fmt.Println(checked)
	} else {
		fmt.Println("Matched successfully")
	}

	///EXAMPLE OUTPUT //////////////////////////////////////////////////////
	// $2a$14$uCMAEUrVLC6KHb1JLzdZnuzpRb4Mp.8nUoQXcKHUNQsC5d5Nwv5qG
	// crypto/bcrypt: hashedPassword is not the hash of the given password
	// Matched successfully
	////////////////////////////////////////////////////////////////////////
}

func main() {
	Example()
}
