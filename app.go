package main

import (
	"github.com/emacsist/alfred3/utils"
	"os"
	"log"
	"bufio"
	"strings"
)

func main() {
	query := utils.GetQuery()

	alfredResponse := utils.NewAlfredResponse()

	lines := Get(query)
	for _, line := range lines {
		data := strings.Fields(line)
		if len(data) < 2 {
			continue
		}
		alfredResponse.AddDefaultItem(data[0] + " => " + strings.Join(data[1:]," "))
	}

	alfredResponse.WriteOutput()
}

func Get(startWith string) []string {
	var data []string
	file, err := os.Open("./http_status.txt")
	if err != nil {
		data = append(data, err.Error())
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(startWith)) > 0 {
			if strings.HasPrefix(line, startWith) {
				data = append(data, scanner.Text())
			}
		} else {
			data = append(data, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		data = append(data, err.Error())
	}
	return data
}
