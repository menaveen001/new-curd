package controller

import "testing"

func TestInsertUser(t *testing.T) {
	got := result
	want := student

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
