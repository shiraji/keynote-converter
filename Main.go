package main

import (
	"os/exec"
	"strings"
	"flag"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

type HtmlSlideList struct {
	Slides []string `json:"slideList"`
}

type HtmlSlide struct {
	Path string
	IsMovie bool
}

func main() {
	flag.Parse()
	htmlDir := flag.Arg(0)
	separator := string(filepath.Separator)
	jsonPath := htmlDir + separator + "assets" + separator + "header.json"
	bytes, _ := ioutil.ReadFile(jsonPath)
	var htmlSlideList HtmlSlideList
	json.Unmarshal(bytes, &htmlSlideList)

	htmlSlides := make([]HtmlSlide, len(htmlSlideList.Slides))
	for index, value := range htmlSlideList.Slides {
		htmlSlides[index] = HtmlSlide{
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
		targetPath := filepath.Dir(filepath.Dir(value))
		targetDir := filepath.Base(targetPath)
		for index, slide := range htmlSlides {
			if slide.Path == targetDir {
				htmlSlides[index].IsMovie = true
			}
		}
	}

}
