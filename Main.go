package main

import (
	"os/exec"
	"strings"
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)

	fmt.Println(dir)

	out, _ := exec.Command("find",
		dir,
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
