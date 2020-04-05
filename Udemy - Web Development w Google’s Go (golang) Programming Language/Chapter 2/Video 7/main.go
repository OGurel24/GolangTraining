package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]

	fmt.Println(os.Args[0])

	html :=
		`<!DOCTYPE html>
<html>
<body>

<h1>` + name + `</h1>

<p>My first paragraph.</p>

</body>
</html>
`

	fmt.Println(html)
}
