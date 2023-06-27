package app

import (
	"fmt"
	"github.com/daniial79/quiz-game/src/quiz"
)

func Start() {
	//1 read and parse problems.csv file
	records := quiz.ReadProblemsFile()
	problems := quiz.ParseProblems(records)
	fmt.Printf("%#v\n", problems)
}
