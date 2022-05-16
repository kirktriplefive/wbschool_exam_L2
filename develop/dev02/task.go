package dev02

//Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
//"a4bc2d5e" => "aaaabccddddde"
//"abcd" => "abcd"
//"45" => "" (некорректная строка)
//"" => ""

import (
	"strconv"
	"strings"
	"unicode"
)

func unpack(str string) string {
	var builder strings.Builder
	reader := strings.NewReader(str)
	char, _, _ := reader.ReadRune()
	if unicode.IsDigit(char) {
		return ""
	}
	for {
		currentChar, _, errForRead := reader.ReadRune(); 
		if errForRead != nil {
			builder.WriteRune(char)
			break
		}

		digit, err := strconv.Atoi(string(currentChar))
		if err == nil {
			builder.WriteString(strings.Repeat(string(char), digit))
		} else {
			builder.WriteRune(char)
		}

		if string(char) == `\` || err == nil {
			char, _, err = reader.ReadRune()
			if errForRead != nil {
				break
			}
		} else {
			char = currentChar
		}
	}
	return builder.String()
}

// func main() {
// 	str:="a4bc2d5e"
// 	fmt.Println(str)
// 	str=unpack(str)
// 	fmt.Println(str)

// }