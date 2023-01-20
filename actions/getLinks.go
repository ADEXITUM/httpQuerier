package checker

import (
	"bufio"
	"fmt"
	"os"
)

type Config struct {
	Name string `json:"name"`
	Link string `json:"string"`
}

func LoadConfig(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var links []string
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}
	return links, nil
}

func GetLinks() []string {
	links, err := LoadConfig("./config/config.dat")
	if err != nil {
		fmt.Println("can't get links")
	}
	return links
}
