package main

import (
	"testing"
)

func Test(t *testing.T) {
	testStr := "jvhvhj:uvhvy:bvhv:jhvv:jhv\nbhjbj:hjbj:jhbjh:jhvj:jhvj\nhvhbj	nkbbb	jbjb	kbjh\nvhbjnk:jhbj	jhbjhb	hbjb\nhvhgvbu:jhbjb"
	want := []string{"uvhvy", "hjbj", "jhbj	jhbjhb	hbjb", "jhbjb"}
	got := cutFunc(testStr, 2, ":")
	i := 0
	for _, str := range got {
		if str != want[i] {
			t.Errorf("Hi() = %q, want %q", str, want[i])
		}
		i++
	}

}

func TestForS(t *testing.T) {
	testStr := "jvhvhj:uvhvy:bvhv:jhvv:jhv\nbhjbj:hjbj:jhbjh:jhvj:jhvj\nhvhbj	nkbbb	jbjb	kbjh\nvhbjnk:jhbj	jhbjhb	hbjb\nhvhgvbu:jhbjb"
	want := []string{"hvhbj	nkbbb	jbjb	kbjh"}
	got := cutFuncWithoutDelimiter(testStr, ":")
	i := 0
	for _, str := range got {
		if str != want[i] {
			t.Errorf("Hi() = %q, want %q", str, want[i])
		}
		i++
	}

}
