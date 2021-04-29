package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("/Users/yangqiang/Downloads/testfile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		datas := strings.Split(scanner.Text(), "|")
		if len(datas) != 40 {
			fmt.Println(datas)
		}
		fmt.Println(len(datas))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}