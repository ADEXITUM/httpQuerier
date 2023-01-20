package main

import (
	"fmt"

	checker "httpQuerier/actions"
)

func main() {
	links := checker.GetLinks()

	output := make([][]string, 0)

	for _, link := range links {
		ch := make(chan string)
		go checker.MakeQuery(link, ch)
		output = append(output, []string{<-ch})
	}

	for _, values := range output {
		fmt.Print(values)
	}
}
