package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func cutFunc(str string, f int, d string) []string {
	strArray := strings.Split(str, "\n")
	resultArray := make([]string, 0)
	arrayOfArray := make([][]string, len(strArray))
	for i, s := range strArray {
		arrayOfArray[i] = strings.Split(s, d)
	}
	for _, s := range arrayOfArray {
		if f-1 < len(s) {
			resultArray = append(resultArray, s[f-1])
		}
	}
	return resultArray

}

func cutFuncWithoutDelimiter(str, d string) []string {
	strArray := strings.Split(str, "\n")
	resultArray := make([]string, 0)
	for _, s := range strArray {
		if !strings.Contains(s, d) {
			resultArray = append(resultArray, s)
		}
	}
	return resultArray

}

func main() {
	var f int
	var d string
	flag.IntVar(&f, "f", 1, "выбрать поля (колонки)")
	flag.StringVar(&d, "d", "\t", "использовать другой разделитель(по умолчанию TAB)")
	useSeparated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строки, для окончания ввода напишите s")
	ftext, err := scanner.ReadString('s')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result: ")
	if *useSeparated {
		s := cutFunc(ftext, f, d)
		for _, str := range s {
			fmt.Println(str)
		}
	} else {
		s := cutFuncWithoutDelimiter(ftext, d)
		for _, str := range s {
			fmt.Println(str)
		}
	}

}
