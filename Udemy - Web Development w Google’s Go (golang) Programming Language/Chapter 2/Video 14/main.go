package main

import (
	"os"
	"text/template"
)

func main() {
	repoRoot, _ := os.Getwd()
	//dir := repoRoot + "/Chapter 2/Video 14"
	tpl, _ := template.ParseGlob(repoRoot + "/onur.gohtml")

	team := []string{"Muslera", "Falcao", "Luyindama", "Onyekuru"}

	tpl.Execute(os.Stdout, team)
}
