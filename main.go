package main

import (
	"fmt"
	"time"

	checker "httpQuerier/actions"
)

func main() {
	tick := time.Tick(1 * time.Minute)
	links := checker.GetLinks()

	output := make([][]string, 0)
	for range tick {
		for _, link := range links {
			ch := make(chan string)
			go checker.MakeQuery(link, ch)
			output = append(output, []string{<-ch})
		}

		for _, values := range output {
			fmt.Print(values)
		}
	}
}
