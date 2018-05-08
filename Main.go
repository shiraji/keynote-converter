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

type SlidePath struct {
	Path string
	IsMovie bool
}

func padNumberWithZero(value int) string {
	return fmt.Sprintf("%03d", value)
}

func main() {
	flag.Parse()
	htmlDir := flag.Arg(0)
	jpegDir := flag.Arg(1)
	jpegFilePrefix := filepath.Base(jpegDir)
	separator := string(filepath.Separator)
	jsonPath := htmlDir + separator + "assets" + separator + "header.json"
	bytes, _ := ioutil.ReadFile(jsonPath)
	var htmlSlideList HtmlSlideList
	json.Unmarshal(bytes, &htmlSlideList)

	slidePaths := make([]SlidePath, len(htmlSlideList.Slides))
	for index, _ := range htmlSlideList.Slides {
		slidePaths[index] = SlidePath{
			jpegDir + separator + jpegFilePrefix + "." + padNumberWithZero(index) + ".jpeg", false,
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
