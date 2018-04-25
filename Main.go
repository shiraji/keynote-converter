package main

import (
	"os/exec"
	"fmt"
	"strings"
	"flag"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)
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
		fmt.Println("-----")
		fmt.Println(value)
	}
}
