package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	want := map[int][]string{
		2: {"ccc", "jjk", "ol"},
		3: {"ррр", "xbbux", "jjk"},
		6: {"jjk"},
	}
	data, _ := ioutil.ReadFile("dev05.txt")
	strs := strings.Split(string(data), "\n")
	got, _, _ := grepFunc(strs, 0, 0, 0, "jjk")
	for key, strArrays := range got {
		for i, str := range strArrays {
			if str != want[key][i] {
				t.Errorf("Hi() = %q, want %q", got, want)
			}
		}
	}

}
