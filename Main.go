package main

import (
	"os/exec"
	"strings"
	"flag"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type SlideList struct {
	Slides []string `json:"slideList"`
}

func main() {
	flag.Parse()
	htmlDir := flag.Arg(0)
	separator := string(filepath.Separator)
	jsonPath := htmlDir + separator + "assets" + separator + "header.json"
	bytes, _ := ioutil.ReadFile(jsonPath)
	var slideList SlideList
	json.Unmarshal(bytes, &slideList)

	for key, slide := range slideList.Slides {
		fmt.Println(strconv.Itoa(key) + ": " + slide)
	}

	fmt.Println(htmlDir)

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
		base := filepath.Base(value)
		targetPath := filepath.Dir(filepath.Dir(value))
		targetDir := filepath.Base(targetPath)
		fmt.Println(targetDir)
		fmt.Println(base)
	}
}
