package dev01

import (
	"testing"
	"time"
)

func TestForError(t *testing.T) {
	var want error
    want = nil
    if got := currentTime(); got != want {
        t.Errorf("Hi() = %q, want %q", got, want)
    }
}

func TestForTime(t *testing.T) {
	want := time.Now()
	_, got := currentTimeWithTime()
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	tim, _ := time.Parse(longForm, want.GoString())
	c, _:= time.Parse(longForm, got.GoString())
	if c != tim {
		t.Errorf("Have = %q, want %q", tim, c)
	}
}