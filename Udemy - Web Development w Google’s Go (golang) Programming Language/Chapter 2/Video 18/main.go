package main

import (
	"os"
	"text/template"
)

type player struct {
	Name        string
	Position    string
	Age         int
	IsImportant bool
}

func main() {
	repoRoot, _ := os.Getwd()
	//dir := repoRoot + "/Chapter 2/Video 18"

	team := []player{{"Muslera", "GK", 32, true},
		{"Falcao", "ST", 35, true}, {"OnurGurel", "RW", 25, false},
		{"Onyekuru", "LW", 24, true}}

	tp := template.Must(template.ParseGlob(repoRoot + "/*.gohtml"))
	tp.Execute(os.Stdout, team)
}
