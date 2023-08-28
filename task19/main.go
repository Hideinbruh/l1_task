package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	word := "Черепаха"
	reversedString := reverse(word)
	fmt.Println(reversedString)
}

func reverse(word string) string {
	dr := []rune(word)
	for i, j := 0, len(dr)-1; i < j; i, j = i+1, j-1 {
		dr[i], dr[j] = dr[j], dr[i]
	}
	return string(dr)
}
