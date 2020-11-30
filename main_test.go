package main

import (
	"log"
	"os"
	"testing"
)

func TestGetTags(t *testing.T) {
	f := "tmp/output.mp3"
	got, _ := GetTags(f)
	want := Song{Filename: "output.mp3", Artist: "Band", Album: "Debut Album", Title: "Sample"}

	if got != want {
		t.Errorf("got %q want %q\n", got, want)
	}

}
func TestMoveFile(t *testing.T) {
	f := "tmp/output.mp3"
	s, _ := GetTags(f)

	err := CreateAlbumDir(s, "tmp")
	if err != nil {
		log.Fatal(err)
	}

	err = MoveFile(s, "tmp")
	if err != nil {
		t.Errorf("move failed, %v", err)
	}
	err = CleanUp(t)
	if err != nil {
		t.Errorf("Cleanup failed, %v", err)
	}
}

func CleanUp(t *testing.T) error {
	f := "tmp/Band/Debut Album/output.mp3"
	stats, err := os.Stat(f)
	if err != nil {
		t.Errorf("not a file, %v", err)
	}

	t.Logf("Name: %s\nPerms: %s\nSize: %d\n", stats.Name(), stats.Mode(), stats.Size())
	err = os.Rename(f, "tmp/output.mp3")
	if err != nil {
		t.Errorf("failed to reset output.mp3, %v", err)
		return err
	}

	err = os.RemoveAll("tmp/Band")
	if err != nil {
		t.Errorf("failed to cleanup directories, %v", err)
		return err
	}
	return nil
}
