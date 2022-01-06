package adventofcode21

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readIntFile(pathToFile string) []int {
	var file *os.File
	var err error
	file, err = os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
