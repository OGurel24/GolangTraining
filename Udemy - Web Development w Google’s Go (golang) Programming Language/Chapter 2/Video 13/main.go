package main

import (
	"os"
	"text/template"
)

func twoTimes(x int) int {
	return x * 2
}

func threeTimes(x int) int {
	return x * 3
}

var fm = template.FuncMap{
	"twoTimes":   twoTimes,
	"threeTimes": threeTimes,
}

func main() {
	repoRoot, _ := os.Getwd()
	dir := repoRoot + "/Chapter 2/Video 13"
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles(dir + "/onur.gohtml"))

	tpl.ExecuteTemplate(os.Stdout, "onur.gohtml", 1)
}
