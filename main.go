package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

//
type QAPair struct {
	Quest string
	Ans   string
}

// set the flags as global variables
var (
	csvFlag = flag.String("csvFile", "problems.csv", "CSV file question,answer")

	timeFlag = flag.Int("timeLimit", 30, "Time limit for the math Quiz")
)

func main() {
	//parsing the flag
	flag.Parse()
	fmt.Println(" Press Enter to Begin")
	var beginTimer string
	fmt.Scanln(&beginTimer)
	problems := readProblems()
	arrStruct := parseProblems(problems)
	compAns(arrStruct)
}

func readProblems() [][]string {

	//reading in the CSV file
	csvFile, err := os.Open(*csvFlag)
	if err != nil {
		fmt.Printf("Your file %s failed to open", *csvFlag)
		os.Exit(1)
	}
	reader := csv.NewReader(csvFile)

	//parsing through each line on the CSV file
	problems, _ := reader.ReadAll()
	if err == io.EOF {
		fmt.Println("Failed to read csv file:", err)
		return nil
	}

	return problems
}

//takes in lines of a 2d string type
// take list of lists and create a list of structs
func parseProblems(problems [][]string) []QAPair {
	//make an array of QA pair

	arrStruct := make([]QAPair, len(problems))
	for i, problem := range problems {
		arrStruct[i] = QAPair{
			Quest: problem[0],
			Ans:   problem[1],
		}

	}
	return arrStruct
}

func compAns(arrStruct []QAPair) {
	correctAns := 0
	incorrectAns := 0

	timeoutChannel := make(chan bool, 1)
	nameChannel := make(chan int, 1)
	userAnswer := ""
	//Calling NewTimer method

	//c := make(chan int)
	for _, questAnsPair := range arrStruct {
		fmt.Println(questAnsPair.Quest)

		go func() {
			fmt.Scanln(&userAnswer)
			nameChannel <- 1
		}()

		go func() {
			time.Sleep(time.Duration(*timeFlag) * time.Second)
			timeoutChannel <- true
		}()

		select {
		case <-nameChannel:

			if userAnswer == questAnsPair.Ans {
				correctAns = correctAns + 1
			} else {
				incorrectAns = incorrectAns + 1
			}
		case <-timeoutChannel:
			fmt.Println("You got", correctAns, "answers correct!")
			fmt.Println("You got", incorrectAns, "answers incorrect")
			fmt.Println("There were", correctAns+incorrectAns, "questions total")
			os.Exit(1)
		}
	}
}
