package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-vgo/robotgo"
)

const INPUT_FILE string = "input.txt"
const DELAY int = 5

func main() {

	list, err := readLines(INPUT_FILE)
	if err != nil {
		fmt.Println(err)
		robotgo.Sleep(DELAY)
		os.Exit(0)
	}

	robotgo.Click()
	robotgo.Sleep(1)
	for _, item := range list {
		fmt.Println(item)
		sendMess(item)
	}
}

func sendMess(mess string) {
	robotgo.TypeStr(mess)
	robotgo.Sleep(1)

	robotgo.KeyTap("enter")
	robotgo.Sleep(1)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := strings.TrimSpace(scanner.Text())
		if content != "" {
			lines = append(lines, content)
		}
	}
	return lines, scanner.Err()
}
