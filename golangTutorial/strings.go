package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Hello World"
	fmt.Println(name)
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")
	printCharsAndBytes(name)

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)

	byteSlice = []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str = string(byteSlice)
	fmt.Println(str)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

	word1 := "Señor"
	length(word1)
	fmt.Println(len(word1))
	word2 := "Pets"
	length(word2)
	fmt.Println(len(word2))

	h := "hello"
	fmt.Println(mutate([]rune(h)))
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

// rune represents a Unicode code point in Go.
// It does not matter how many bytes the code point occupies, it can be represented by a rune.
func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}
