package main

import (
	"os"
	"text/template"
)

type player struct {
	name     string
	position string
	age      int64
}

func main() {
	repoRoot, _ := os.Getwd()
	dir := repoRoot + "/Chapter 2/Video 11"
	tpl, _ := template.ParseGlob(dir + "/onur.gohtml")

	team := []player{{"Muslera", "GK", 32}, {"Falcao", "ST", 35}}

	tpl.Execute(os.Stdout, team)
}
