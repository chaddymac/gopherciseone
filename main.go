package main

import (
	"encoding/csv"
	"flag"
	fmt "fmt"
	_ "go/parser"
	"io"
	"os"
)

//
type QAPair struct {
	Quest string
	Ans   string
}

func main() {

	problems := readProblems()
	arrStruct := parseProblems(problems)
	compAns(arrStruct)
}

func readProblems() [][]string {

	//reading in the CSV file
	csvFlag := flag.String("csvFile", "problems.csv", "CSV file question,answer")
	flag.Parse()

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
	for _, questAnsPair := range arrStruct {
		fmt.Println(questAnsPair.Quest)
		var userAnswer string
		fmt.Scanln(&userAnswer)
		if userAnswer == questAnsPair.Ans {
			correctAns = correctAns + 1
		} else {
			incorrectAns = incorrectAns + 1
		}

	}
	fmt.Println("You got", correctAns, "answers correct!")
	fmt.Println("You got", incorrectAns, "answers incorrect")
	fmt.Println("There were", correctAns+incorrectAns, "questions total")
}
