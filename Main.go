package main

import (
	"os/exec"
	"strings"
	"flag"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type SlideList struct {
	Slides []string `json:"slideList"`
}

type Slide struct {
	Path string
	IsMovie bool
}

func main() {
	flag.Parse()
	htmlDir := flag.Arg(0)
	separator := string(filepath.Separator)
	jsonPath := htmlDir + separator + "assets" + separator + "header.json"
	bytes, _ := ioutil.ReadFile(jsonPath)
	var slideList SlideList
	json.Unmarshal(bytes, &slideList)

	slides := make([]Slide, len(slideList.Slides))
	for index, value := range slideList.Slides {
		slides[index] = Slide {
			value, false,
		}
	}

	out, _ := exec.Command("find",
		htmlDir,
		"-not", "-name", "*.pdf*",
		"-not", "-name", "*.json*",
		"-not", "-name", "*.jpeg",
		"-not", "-name", "*.png",
		"-not", "-name", "*.js",
		"-not", "-name", "*.bcmap",
		"-not", "-name", "*.DS_Store",
		"-not", "-name", "*.html",
		"-not", "-name", "*.css",
		"-not", "-name", "LICENSE",
		"-type", "f",
	).Output()

	results := strings.Split(string(out), "\n")

	for _, value := range results {
		//base := filepath.Base(value)
		targetPath := filepath.Dir(filepath.Dir(value))
		targetDir := filepath.Base(targetPath)
		for index, slide := range slides {
			if slide.Path == targetDir {
				slides[index].IsMovie = true
			}
		}
	}

	for _, v := range slides {
		fmt.Println(v)
	}
}
