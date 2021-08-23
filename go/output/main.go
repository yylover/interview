package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)


func getTodayUnix() int64 {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime.Unix()
}


func main() {

	fmt.Println(getTodayUnix())
	return ;

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