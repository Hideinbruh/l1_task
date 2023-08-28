package main

import "fmt"

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	words := []string{"apple", "this", "eats", "he"}
	reverseWords := make([]string, 0)
	for i := 0; i < len(words); i++ {
		reverseWords = append(reverseWords, words[len(words)-i-1])
	}
	fmt.Println(reverseWords)
}
