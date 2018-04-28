package main

import (
	"fmt"

	"github.com/prasanjit/udemyTraining/04_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("hello"))
	fmt.Println(stringutil.Reverse(stringutil.MyName))
}
