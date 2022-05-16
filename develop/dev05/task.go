package main

import (
	"flag"
	"fmt"
)

func main() {
	var n int
	flag.IntVar(&n, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&n, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&n, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	useCount := flag.Bool("c", false, "количество строк")
	useIgnore := flag.Bool("i", false, "вместо совпадения, исключать")
	useFixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	useLineNum := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()
	if str := flag.Arg(0); str != "" {
		
	} else {
		fmt.Println("Вы не ввели строку для фильтрации!")
	}
}