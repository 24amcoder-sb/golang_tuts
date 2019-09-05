package main

import (
	"fmt"
	"math"

	"github.com/24amcoder-sb/golang_tuts/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(32.7))
	fmt.Println(math.Ceil(32.7))
	fmt.Println(math.Sqrt(49))

	fmt.Println(strutil.Reverse("muspI meroL"))
}
