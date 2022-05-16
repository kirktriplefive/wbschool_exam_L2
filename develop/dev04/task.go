package main

import (
	"fmt"
	"sort"
	"strings"
)

func getAnagrams(arr *[]string) *map[string][]string {
	result := make(map[string][]string)
	testArr := *arr
	//fmt.Println(len(testArr))
	for i, s := range testArr {
		result[s] = append(result[s], s)
		r := len(testArr)
		for j := i + 1; j <= r-1; {
			f, num := isAnagram(s, testArr[j])
			testArr[j] = strings.ToLower(testArr[j])
			if f && num == 2 {
				result[s] = append(result[s], testArr[j])
				copy(testArr[j:], testArr[j+1:])
				testArr[len(testArr)-1] = ""
				testArr = testArr[:len(testArr)-1]
			} else if f && num == 1 {
				copy(testArr[j:], testArr[j+1:])
				testArr[len(testArr)-1] = ""
				testArr = testArr[:len(testArr)-1]
			} else if !f {
				j++
			}
			r = len(testArr)
		}
	}
	for key := range result {
		if len(result[key]) == 1 || key == "" {
			delete(result, key)
		}
		sort.Strings(result[key])
	}
	return &result

}

//main
func main() {
	array := []string{"пятак", "Слиток", "пятак", "столик", "тяпка", "пятка", "кирилл", "листок"}
	result := getAnagrams(&array)
	fmt.Println(*result)
}

func isAnagram(original, test string) (bool, int) {
	// СРавниваем строки без пробелов по бокам. Если равны, то они не явлются анаграммами
	if strings.TrimSpace(original) == strings.TrimSpace(test) {
		return true, 1
	}

	// Удаляем все пробелы в строках. Делаем нижний регистр для
	// строки и сортируем для сравнения
	strOut := sortString(strings.ToLower(strings.Join(strings.Fields(original), "")))
	testOut := sortString(strings.ToLower(strings.Join(strings.Fields(test), "")))
	// Сравниваем отсортированные строки, если True, то это анаграмма
	if strOut == testOut {
		//fmt.Println(strOut, ""+testOut)
		return true, 2
	} else {
		return false, 0
	}
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
