package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"flag"
	"strings"
	"time"
)

func parseQuestionsFromCSV(records [][]string) []problem {
	problems := make([]problem, len(records))
	for i, record := range records {
		problems[i] = problem{
			question: record[0],
			solution: strings.TrimSpace(record[1]),
		}
	}
	return problems
}

type problem struct {
	question string
	solution string
}

func main() {
	fileName := flag.String("csv", "problems.csv", "the name of the csv file consisting of problems")
	timeLimit := flag.Int("time", 20, "time limit for the quiz")
	flag.Parse()


	csvfile, err := os.Open(*fileName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	problems := parseQuestionsFromCSV(records)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	numCorrect := 0

	for _, problem := range problems {
		fmt.Println(problem.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <- timer.C:
			fmt.Printf("Score: %d/%d", numCorrect, len(problems))
			return
		case answer := <- answerChannel:
			if answer == problem.solution {
				numCorrect++
			}
		}
			
	}

	fmt.Printf("Score: %d/%d", numCorrect, len(problems))
}
