package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func IsUnique(str string) bool {
	newStr := strings.ToLower(str)
	cr := make(map[rune]int)

	for _, elem := range newStr {
		cr[elem]++

		if cr[elem] > 1 {
			return false
		}
	}
	return true
}

func main() {
	str1 := "abdc"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(IsUnique(str1))
	fmt.Println(IsUnique(str2))
	fmt.Println(IsUnique(str3))
}
