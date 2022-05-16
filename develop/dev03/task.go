package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func defaultSort(strs []string) []string {
	sort.Strings(strs)
	return strs

}

func defaultSortByNumber(strs []string) ([]string, error) {
	strArrays := make([][]string, len(strs))
	intArrays := make([]int, len(strs))
	var result []string
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
	return result, nil

}

func sortByColumn(strs []string, k int) []string {
	k = k - 1
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
	sort.Strings(strs)
	for i, s := range strs {
		strArrays[i] = strings.Split(s, " ")
	}
	for _, s := range strArrays {
		s[0], s[k] = s[k], s[0]
	}
	for i, s := range strArrays {
		strs[i] = strings.Join(s, " ")
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
		strs[i], strs[len(strs)-i-1] = strs[len(strs)-i-1], strs[i]
	}
	return strs
}

func deleteSort(strs []string) []string {
	j := 1
	for i := 1; i < len(strs); i++ {
		if strs[i] != strs[i-1] {
			strs[j] = strs[i]
			j++
		}
	}
	return strs[:j]
}

func main() {
	var k int
	var dump string
	var result []string

	flag.IntVar(&k, "k", 0, "The integer param")
	flag.StringVar(&dump, "dump", "default_dump", "The name of a dump")
	useNumeric := flag.Bool("n", false, "сортировать по числовому значению")
	useBack := flag.Bool("r", false, "сортировать в обратном порядке")
	useUnimaginable := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	if filename := flag.Arg(0); filename != "" {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			if *useNumeric {
				if k != 0 {
					strs := strings.Split(string(data), "\n")
					if result, err = sortByColumnByNumber(strs, k); err != nil {
						fmt.Println(err)
					}
				} else {
					strs := strings.Split(string(data), "\n")
					if result, err = defaultSortByNumber(strs); err != nil {
						fmt.Println(err)
					}
				}
			} else {
				if k != 0 {
					strs := strings.Split(string(data), "\n")
					result = sortByColumn(strs, k)
				} else {
					strs := strings.Split(string(data), "\n")
					result = defaultSort(strs)
				}
			}

			if *useBack {
				result = reverseSort(result)
			}

			if *useUnimaginable {
				result = deleteSort(result)
			}

			for _, s := range result {
				fmt.Println(s)
			}

		}
	}

}
