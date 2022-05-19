package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Отсортировать строки в файле по аналогии с консольной утилитой sort 
// (man sort — смотрим описание и основные параметры): на входе подается файл из 
// несортированными строками, на выходе — файл с отсортированными.


func main() {
	var k int
	var dump string
	var result []string

	flag.IntVar(&k, "k", 0, "The integer param")
	flag.StringVar(&dump, "dump", "default_dump", "The name of a dump")
	useNumeric := flag.Bool("n", false, "сортировать по числовому значению")
	useBack := flag.Bool("r", false, "сортировать в обратном порядке")
	useUnimaginable := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()//парсим флаги
	if filename := flag.Arg(0); filename != "" {//считываем имя файла
		data, err := ioutil.ReadFile(filename)//читаем из файла в data 
		if err != nil {
			fmt.Println(err)
		} else {
			if *useNumeric {//если сортировка нужна по числу
				if k != 0 { //если нужно сортировать по определенной колонке
					strs := strings.Split(string(data), "\n")//разбиваем строку на массив строк с разделителем ENTER
					if result, err = sortByColumnByNumber(strs, k); err != nil {// вызываем функцию для сортировки по числовому значению
						fmt.Println(err)
					}
				} else {
					strs := strings.Split(string(data), "\n")//сортировка не по колонке (дефолтная - по первому числу)
					if result, err = defaultSortByNumber(strs); err != nil {
						fmt.Println(err)
					}
				}
			} else { //если не по числу
				if k != 0 {
					strs := strings.Split(string(data), "\n")
					result = sortByColumn(strs, k)//сортируем по колнке
				} else {
					strs := strings.Split(string(data), "\n")
					result = defaultSort(strs)// сортируем по дефолту
				}
			}

			if *useBack {
				result = reverseSort(result)//реверс сортировки 
			}

			if *useUnimaginable {
				result = deleteSort(result)// не выводим повторяющиеся строки (удаляем их из результата)
			}

			for _, s := range result {
				fmt.Println(s)
			}

		}
	}

}

func defaultSort(strs []string) []string { //функция дефолтной сортировки
	sort.Strings(strs)
	return strs

}

func defaultSortByNumber(strs []string) ([]string, error) { //функция дефолтной сортировки по числовому значению
	strArrays := make([][]string, len(strs))
	intArrays := make([]int, len(strs))
	var result []string
	for i, s := range strs {
		strArrays[i] = strings.Split(s, " ")//разбиваем массив строк на массив массивов (по словам)
	}
	for i, s := range strArrays {
		if j, err := strconv.Atoi(s[0]); err != nil { // берем числовое значение первого слова (если можно)
			return nil, err
		} else {
			intArrays[i] = j//запоминаем его в массив
		}
	}
	sort.Ints(intArrays)//сортируем массив числовых значений
	for _, s := range intArrays {
		for _, str := range strArrays {
			if first, err := strconv.Atoi(str[0]); err == nil {
				if s == first {//далее добавляем в массив те строки, которым соответсвуют наши остортированные числовые значения
					resString := strings.Join(str, " ")
					result = append(result, resString)
				}

			}

		}
	}
	return result, nil

}

func sortByColumn(strs []string, k int) []string {
	k = k - 1//номер колонки на 1 меньше, так как первый элемент массива - нулеывой
	strArrays := make([][]string, len(strs))
	for i, s := range strs {
		strArrays[i] = strings.Split(s, " ")//разбиваем на массив массивов
	}
	for _, s := range strArrays {
		s[0], s[k] = s[k], s[0]//меняем нужную колонку с первой местами
	}
	for i, s := range strArrays {
		strs[i] = strings.Join(s, " ")//соединяем в массив строк
	}
	sort.Strings(strs)//сортируем
	for i, s := range strs {
		strArrays[i] = strings.Split(s, " ")//массив массивов
	}
	for _, s := range strArrays {
		s[0], s[k] = s[k], s[0]//меняем обратно
	}
	for i, s := range strArrays {
		strs[i] = strings.Join(s, " ")//соединяем
	}
	return strs
}

func sortByColumnByNumber(strs []string, k int) ([]string, error) {
	k = k - 1
	var result []string
	intArrays := make([]int, len(strs))
	strArrays := make([][]string, len(strs))
	for i, s := range strs {
		strArrays[i] = strings.Split(s, " ")
	}
	for _, s := range strArrays {
		s[0], s[k] = s[k], s[0]
	}
	for i, s := range strArrays {
		strs[i] = strings.Join(s, " ")
	}
	for i, s := range strs {
		strArrays[i] = strings.Split(s, " ")
	}
	for i, s := range strArrays {
		if j, err := strconv.Atoi(s[0]); err != nil {
			return nil, err
		} else {
			intArrays[i] = j
		}
	}
	sort.Ints(intArrays)
	for _, s := range intArrays {
		for _, str := range strArrays {
			if first, err := strconv.Atoi(str[0]); err == nil {
				if s == first {
					resString := strings.Join(str, " ")
					result = append(result, resString)
				}

			}

		}
	}
	for i, s := range result {
		strArrays[i] = strings.Split(s, " ")
	}
	for _, s := range strArrays {
		s[0], s[k] = s[k], s[0]
	}
	for i, s := range strArrays {
		strs[i] = strings.Join(s, " ")
	}
	return strs, nil
}

func reverseSort(strs []string) []string {
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-i-1] = strs[len(strs)-i-1], strs[i]//меняем порядок на обратный
	}
	return strs
}

func deleteSort(strs []string) []string {
	j := 1
	for i := 1; i < len(strs); i++ {//удаляем  повторяющиеся
		if strs[i] != strs[i-1] {
			strs[j] = strs[i]
			j++
		}
	}
	return strs[:j]
}

