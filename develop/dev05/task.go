package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

// Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).


func main() {
	var a int
	var b int
	var c int
	flag.IntVar(&a, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&b, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&c, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	useCount := flag.Bool("c", false, "количество строк")
	useIgnore := flag.Bool("i", false, "вместо совпадения, исключать")
	useFixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	useLineNum := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()// парсим флаги
	if str := flag.Arg(0); str != "" {// считываем строку для работы
		if filename := flag.Arg(1); filename != "" {// берем имя файла
			data, err := ioutil.ReadFile(filename) // читаем файл
			strs := strings.Split(string(data), "\n") // записываем строки из файла в массив
			if err != nil {
				fmt.Println(err)
			} else {
				if *useFixed {// если нужно точное совпадение со строкой
					r := grepFuncUseFixed(strs, str)
					for _, res := range r {
						fmt.Println(res)
					}
				} else { // иначе ищем подстроку в строках
					resNorm, resInvert, err := (grepFunc(strs, a, b, c, str))
					if err != nil {
						fmt.Println(err)
					} else {
						if *useIgnore && *useLineNum {
							for key, str := range resInvert {
								res := strings.Join(str, " ")
								fmt.Println(key+1, res)
							}
						} else if *useIgnore && *useCount {
							fmt.Println(len(resInvert))
						} else if *useIgnore {
							for _, str := range resInvert {
								res := strings.Join(str, " ")
								fmt.Println(res)
							}
						} else if *useCount {
							fmt.Println(len(resNorm))
						} else if *useLineNum {
							for key, str := range resNorm {
								res := strings.Join(str, " ")
								fmt.Println(key+1, res)
							}
						} else {
							for _, str := range resNorm {
								res := strings.Join(str, " ")
								fmt.Println(res)
							}
						}
					}

				}
			}

		}

	} else {
		fmt.Println("Вы не ввели строку для фильтрации!")
	}
}

func grepFuncUseFixed(strs []string, str string) map[int]string {
	result := make(map[int]string)
	for i, st := range strs {
		if st == str {// если полностью совпадает
			result[i] = str
		}
	}
	return result
}

func grepFunc(strs []string, a, b, c int, str string) (map[int][]string, map[int][]string, error) {
	resNum := make([]int, 0)
	resultNorm := make(map[int][]string)
	resultInvert := make(map[int][]string)
	if a != 0 && b != 0 && c != 0 {
		return nil, nil, fmt.Errorf("Вы ввели неверные флаги, флаги a, b, c работают по-отдельности")
	} else if (a != 0 && b != 0) || (a != 0 && c != 0) || (b != 0 && c != 0) {
		return nil, nil, fmt.Errorf("Вы ввели неверные флаги, флаги a, b, c работают по-отдельности")
	} else {
		for i, st := range strs {
			arr := strings.Split(st, " ")
			resultInvert[i] = arr
			for _, s := range arr {
				if s == str {
					resNum = append(resNum, i)
					resultNorm[i] = arr
					delete(resultInvert, i)
				}
			}
		}
		//fmt.Println(resNum)
		if a > 0 {
			for _, i := range resNum {
				if i+a >= len(strs) {
					return nil, nil, fmt.Errorf("Не хватает строк для команды after")
				} else {
					for j := i + 1; j < a+i+1; j++ {
						arr := strings.Split(strs[j], " ")
						resultNorm[j] = arr
						delete(resultInvert, j)
					}
				}
			}
		}

		if b > 0 {
			for _, i := range resNum {
				if i-b < 0 {
					return nil, nil, fmt.Errorf("Не хватает строк для команды before")
				} else {
					for j := i - 1; j > i-b-1; j-- {
						arr := strings.Split(strs[j], " ")
						resultNorm[j] = arr
						delete(resultInvert, j)
					}
				}
			}
		}

		if c > 0 {
			a = c
			b = c
			for _, i := range resNum {
				if i-b < 0 {
					return nil, nil, fmt.Errorf("Не хватает строк для команды context-before")
				} else {
					for j := i - 1; j > i-b-1; j-- {
						arr := strings.Split(strs[j], " ")
						resultNorm[j] = arr
						delete(resultInvert, j)
					}
				}
			}
			for _, i := range resNum {
				if i+a >= len(strs) {
					return nil, nil, fmt.Errorf("Не хватает строк для команды context")
				} else {
					for j := i + 1; j < a+i+1; j++ {
						arr := strings.Split(strs[j], " ")
						resultNorm[j] = arr
						delete(resultInvert, j)
					}
				}
			}

		}

		//fmt.Println(resultNorm)
		//fmt.Println(resultInvert)
		return resultNorm, resultInvert, nil
	}

}
