package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	ques string
	ans  string
}

func parseRecords(records [][]string) []problem {
	parsedRecords := make([]problem, len(records))
	for i, record := range records {
		parsedRecords[i] = problem{record[0], strings.TrimSpace(record[1])}
	}
	return parsedRecords
}

func readCsvData(path string) []problem {
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
	parsedRecords := parseRecords(records)
	return parsedRecords
}

func quizTimer(sec int, scoreRef *int) {
	t := time.NewTimer(time.Duration(sec) * time.Second)
	<-t.C
	fmt.Println("\n\nTimes up!!!")
	fmt.Println("Your score = ", *scoreRef)
	os.Exit(0)
}

func conductQuiz(records []problem, duration int) int {
	score := 0
	var sol string

	go quizTimer(duration, &score)
	for i, record := range records {
		fmt.Printf("Problem #%d: %s \nYour answer:", i, record.ques)
		fmt.Scan(&sol)
		fmt.Println()
		if sol == record.ans {
			score++
		}
	}
	return score
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in ques,ans foramt")
	duration := flag.Int("duration", 10, "duration of the quiz")
	flag.Parse()

	records := readCsvData(*csvFileName)

	score := conductQuiz(records, *duration)
	fmt.Println("Your score is ", score)
}
