package main

import (
	"os"
	"text/template"
)

func main() {
	repoRoot, _ := os.Getwd()
	dir := repoRoot + "/Chapter 2/Video 8"
	tpl, _ := template.ParseGlob(dir + "/onur.gohtml")
	tpl.Execute(os.Stdout, nil)
	tpl, _ = template.ParseGlob(dir + "/*.gohtml")
	tpl.ExecuteTemplate(os.Stdout, "onur2.gohtml", nil)

}
