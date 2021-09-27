package main

import (
	"fmt"
	"math"
	"github.com/martin-stankard/go_crash_course/03_pkgs/strutil"
	//remember to run "go mod init" for this part to work
)

func main() {
	fmt.Println(math.Floor(2.999999))
	fmt.Println(math.Ceil(2.999999))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}
