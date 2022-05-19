package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Реализовать утилиту аналог консольной команды cut (man cut). 
// Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.


func cutFunc(str string, f int, d string) []string {
	strArray := strings.Split(str, "\n") // берем массив строк
	resultArray := make([]string, 0)
	arrayOfArray := make([][]string, len(strArray))
	for i, s := range strArray {
		arrayOfArray[i] = strings.Split(s, d)// разбиваем на массив массивов где разделитель - заданный разделитель
	}
	for _, s := range arrayOfArray {
		if f-1 < len(s) {
			resultArray = append(resultArray, s[f-1])// берем нужную колонку
		}
	}
	return resultArray

}

func cutFuncWithoutDelimiter(str, d string) []string {
	strArray := strings.Split(str, "\n")
	resultArray := make([]string, 0)
	for _, s := range strArray {
		if !strings.Contains(s, d) { // проверяем если нет разделителя в строке то
			resultArray = append(resultArray, s) // добавляем к результату
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
	flag.Parse() // парсим флаги
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строки, для окончания ввода напишите s")
	ftext, err := scanner.ReadString('s')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result: ")
	if *useSeparated { // разбиваем на колонки по разделителю
		s := cutFunc(ftext, f, d)
		for _, str := range s {
			fmt.Println(str)
		}
	} else {
		s := cutFuncWithoutDelimiter(ftext, d) // строки, не содержащие разделитель
		for _, str := range s {
			fmt.Println(str)
		}
	}

}
