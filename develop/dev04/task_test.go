package main

import (
	"testing"
)

func Test(t *testing.T) {
	want := map[string][]string{
		"кирилл": {"ирллик", "кирилл"},
		"нос":    {"нос", "осн", "сон"},
	}
	array := []string{"кирилл", "нос", "сон", "ИрЛлик", "осН", "дом"}
	got := getAnagrams(&array)
	for key, strArrays := range *got {
		for i, str := range strArrays {
			if str != want[key][i] {
				t.Errorf("Hi() = %q, want %q", got, want)
			}
		}
	}

}
