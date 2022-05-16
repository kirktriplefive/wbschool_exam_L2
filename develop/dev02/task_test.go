package dev02

import "testing"

func TestForError(t *testing.T) {
	want:="aaaabccddddde"
	got:="a4bc2d5e"
    if got = unpack(got); got != want {
        t.Errorf("Have = %q, want %q", got, want)
    }

	want="abcd"
	got="abcd"
    if got = unpack(got); got != want {
        t.Errorf("Have = %q, want %q", got, want)
    }

	want=""
	got="45"
    if got = unpack(got); got != want {
        t.Errorf("Have = %q, want %q", got, want)
    }

	want=`qwe\\\\\`
	got=`qwe\\5`
    if got = unpack(got); got != want {
        t.Errorf("Have = %q, want %q", got, want)
    }

    want=`qwe45`
	got=`qwe\4\5`
    if got = unpack(got); got != want {
        t.Errorf("Have = %q, want %q", got, want)
    }

    want=`qwe44444`
	got=`qwe\45`
    if got = unpack(got); got != want {
        t.Errorf("Have = %q, want %q", got, want)
    }
}