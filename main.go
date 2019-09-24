package main

import (
	"flag"
	"github.com/bogem/id3v2"
	"github.com/schollz/progressbar"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Song struct {
	Filename string
	Artist   string
	Album    string
	Title    string
}

func main() {
	directory := flag.String("directory", "", "Path to mp3 files")
	flag.Parse()

	log.Printf("Directory: %s\n", *directory)

	if *directory == "" {
		log.Fatal("Directory required")
		os.Exit(1)
	}

	pattern := *directory + "/" + "*.mp3"

	var Songs []Song

	log.Println("Gathering files")
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.New(len(files))

	log.Println("Reading tags")
	for _, file := range files {
		t, err := GetTags(file)
		if err != nil {
			log.Fatal(err)
		}
		bar.Add(1)
		Songs = append(Songs, t)
	}

	bar = progressbar.New(len(files))
	log.Printf("\nRelocating song\n")
	for _, song := range Songs {
		err := CreateAlbumDir(song, *directory)
		if err != nil {
			log.Fatal(err)
		}
		err = MoveFile(song, *directory)
		if err != nil {
			log.Fatal(err)
		}
		bar.Add(1)
	}
}

func GetTags(f string) (Song, error) {
	s := Song{}

	tag, err := id3v2.Open(f, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error opening mp3 file: ", err)
		return Song{}, err
	}

	if tag.Artist() == "" {
		log.Fatalf("Empty Artist tag found on %s, aborting\n", f)
		os.Exit(1)
	}

	if tag.Album() == "" {
		log.Fatalf("Empty Album tag found on %s, aborting\n", f)
		os.Exit(1)
	}

	defer tag.Close()
	splitFilename := strings.Split(f, "/")
	s.Filename = splitFilename[len(splitFilename)-1]
	s.Artist = tag.Artist()
	s.Title = tag.Title()
	s.Album = tag.Album()

	return s, nil
}

func CreateAlbumDir(s Song, p string) error {
	err := os.MkdirAll(p+"/"+s.Artist+"/"+s.Album, 0755)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func MoveFile(s Song, p string) error {
	err := os.Rename(p+"/"+s.Filename, p+"/"+s.Artist+"/"+s.Album+"/"+s.Filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
