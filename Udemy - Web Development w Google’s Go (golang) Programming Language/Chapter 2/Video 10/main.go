package main

import (
	"os"
	"text/template"
)

func main() {
	repoRoot, _ := os.Getwd()
	dir := repoRoot + "/Chapter 2/Video 10"
	tpl, _ := template.ParseGlob(dir + "/onur.gohtml")
	tpl.Execute(os.Stdout, "Onur")

}
