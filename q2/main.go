package main

import "fmt"

func one(x *int) {
	*x = 0
}

func main() {

	string := "helloworldplayground"

	var alphabet [26]int

	for i := 0; i < len(string); i++ {
		alphabet[string[i]-'a']++
	}

	max := alphabet[0]
	var maxChar byte
	for i := 0; i < len(alphabet); i++ {
		if max < alphabet[i] {
			max = alphabet[i]
			maxChar = byte('a' + i)
		}
	}
	fmt.Printf("%c", maxChar)
}
