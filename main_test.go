package main

import "testing"

func TestGetTags(t *testing.T) {
	f := "tmp/output.mp3"
	got, _ := GetTags(f)
	want := Song{Filename: "output.mp3", Artist: "Band", Album: "Debut Album", Title: "Sample"}

	if got != want {
		t.Errorf("got %q want %q\n", got, want)
	}
}
