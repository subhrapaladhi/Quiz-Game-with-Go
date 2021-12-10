package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsvData(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records, err
}

func conductQuiz(records [][]string) int {
	score := 0
	var sol string
	for i := 0; i < len(records); i++ {
		fmt.Println("#Problem", i, ":", records[i][0])
		fmt.Print("Your answer: ")
		fmt.Scan(&sol)
		fmt.Println()
		if sol == records[i][1] {
			score++
		}
	}
	return score
}

func main() {
	records, err := readCsvData("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	score := conductQuiz(records)
	fmt.Println("Your score is ", score)
}
