package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	words := make([]string, 0)

	for _, arg := range args {
		words = append(words, strings.Fields(arg)...)
	}

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	// キモくなる順に並び替えてある
	// 'i' は大文字にしてもあんまりキモくないので除外
	vowels := []rune{
		'o', 'u', 'e', 'a',
	}
	vowelCounts := make(map[rune]uint8)

	for _, vowel := range vowels {
		vowelCounts[vowel] = 0
	}

	for _, word := range words {
		for _, char := range word {
			_, exists := vowelCounts[char]
			if exists {
				vowelCounts[char] += 1
			}
		}
	}

	mostVowel := vowels[0]
	mostVowelCount := vowelCounts[mostVowel]

	for _, vowel := range vowels {
		count := vowelCounts[vowel]
		if count > mostVowelCount {
			mostVowel = vowel
			mostVowelCount = count
		}
	}

	for i, word := range words {
		strangeWord := ""
		for _, char := range word {
			if char == mostVowel {
				strangeWord += strings.ToUpper(string(char))
			} else {
				strangeWord += string(char)
			}
		}
		words[i] = strangeWord
	}

	strangeName := strings.Join(words, " ")
	fmt.Println(strangeName)
}
