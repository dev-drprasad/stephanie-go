package main

import (
	"fmt"

	"github.com/dev-drprasad/stephanie-go/mentors"
)

func main() {
	result := mentors.ScrapeMentors()
	fmt.Println(result)
}
