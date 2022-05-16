package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	//"time"
)

func TestForError(t *testing.T) {
	var want error
	want = nil
	if data, err := ioutil.ReadFile("dev03.txt"); err != nil {
		fmt.Println(err)
	} else {
		strs := strings.Split(string(data), "\n")
		if _, got := defaultSortByNumber(strs); got != want {
			t.Errorf("Hi() = %q, want %q", got, want)
		}
	}

}

func TestForColoumn(t *testing.T) {
	want := []string{
		"66 aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa ceec 77 wswz ttyn",
		"2234 ggg ecece 12 cece ecec cec",
		"55 ggg ecece 655 cece ecec cec",
		"34 nn giriceic decec ecece",
		"4 b ll 9892 pedx yrc x",
	}
	if data, err := ioutil.ReadFile("dev03.txt"); err != nil {
		fmt.Println(err)
	} else {
		strs := strings.Split(string(data), "\n")
		got := sortByColumn(strs, 3)
		for i, s := range got {
			if s!=want[i] {
				t.Errorf("Hi() = %q, want %q", got, want)
			}
		}

	}
}
