package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()_+{}[]:;<>,.?/~"
)

func fullfilCriteria(charSet string, clen int) string {
	var result string
	for i := 0; i < clen; i++ {
		result = result + string(charSet[rand.Intn(len(charSet))])
		//result = result + string(charSet[rand.Intn(len(charSet))])
	}

	return result
}

func generatePassword(criteria map[string]int, plen int) string {
	rand.Seed(time.Now().UnixNano())

	var password string

	for k, v := range criteria {
		switch k {
		case "uppercaseLetter":
			password = password + fullfilCriteria(uppercaseLetters, v)
		case "lowercaseLetter":
			password = password + fullfilCriteria(lowercaseLetters, v)
		case "digit":
			password = password + fullfilCriteria(digits, v)
		case "specialChar":
			password = password + fullfilCriteria(specialChars, v)
		}
	}

	for i := len(password); i < plen; i++ {
		randomCharSet := rand.Intn(4)
		charSet := ""
		switch randomCharSet {
		case 0:
			charSet = uppercaseLetters
		case 1:
			charSet = lowercaseLetters
		case 2:
			charSet = digits
		case 3:
			charSet = specialChars
		}
		password = password + string(charSet[rand.Intn(len(charSet))])
	}

	return password
}

func generatePasswords(criteria map[string]int, plen int, numPasswords int) []string {
	var result []string
	for i := 0; i < numPasswords; i++ {
		password := generatePassword(criteria, plen)
		result = append(result, password)
	}
	return result
}

func main() {
	var criteria = map[string]int{
		"uppercaseLetter": 3,
		"lowercaseLetter": 3,
		"digit":           2,
		"specialChar":     2,
	}

	numPasswords := 5
	plen := 10
	fmt.Printf("Generated Sample Passwords:\n")
	fmt.Println(generatePasswords(criteria, plen, numPasswords))
}
