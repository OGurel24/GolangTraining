package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	v1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	v2 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	fmt.Println(v1)
	fmt.Println(v2)

}
