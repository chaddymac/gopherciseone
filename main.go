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
	nameChannel := make(chan bool, 1)
	userAnswer := ""

	for _, questAnsPair := range arrStruct {
		fmt.Println(questAnsPair.Quest)
		//go routines allow you to run this code without stopping the program.
		// allows you to start the countdown while waiting for the users input.
		go func() {
			fmt.Scanln(&userAnswer)
			nameChannel <- true
		}()
		//
		go func() {
			time.Sleep(time.Duration(*timeFlag) * time.Second)
			timeoutChannel <- true
		}()
		//select catches the channels when they get set. similar to a javascript onclick
		//channel acts as a listener for the channels
		select {
		case <-nameChannel:

			if userAnswer == questAnsPair.Ans {
				correctAns = correctAns + 1
			} else {
				incorrectAns = incorrectAns + 1
			}
			//prints the following statements once the time is up
		case <-timeoutChannel:
			fmt.Println("You got", correctAns, "answers correct!")
			fmt.Println("You got", incorrectAns, "answers incorrect")
			fmt.Println("There were", correctAns+incorrectAns, "questions total")
			os.Exit(1)
		}
	}
}
