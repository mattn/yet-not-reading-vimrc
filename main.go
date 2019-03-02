package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Entry struct {
	ID     int    `yaml:"id"`
	Date   string `yaml:"date"`
	Author struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	} `yaml:"author"`
	Vimrcs []struct {
		Name   string `yaml:"name"`
		RawURL string `yaml:"raw_url"`
		URL    string `yaml:"url"`
	} `yaml:"vimrcs"`
	Part    interface{} `yaml:"part"`
	Other   interface{} `yaml:"other"`
	Members []string    `yaml:"members"`
	Log     string      `yaml:"log"`
	Links   []string    `yaml:"links,omitempty"`
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	resp, err := http.Get("https://raw.githubusercontent.com/vim-jp/reading-vimrc/gh-pages/_data/archives.yml")
	if err != nil {
		log.Fatal(err)
	}
	var entries []Entry
	err = yaml.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		log.Fatal(err)
	}
	authors := map[string]struct{}{}
	want := os.Args[1]
	for _, entry := range entries {
		name := entry.Author.Name
		if _, ok := authors[want]; ok {
			fmt.Println(name + "'s vimrc is already readed")
			os.Exit(1)
		}
		authors[name] = struct{}{}
	}
	fmt.Println(want + "'s vimrc is not readed yet")
}
