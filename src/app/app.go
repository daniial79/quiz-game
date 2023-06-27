package app

import (
	"fmt"
	"github.com/daniial79/quiz-game/src/quiz"
)

func Start() {

	records := quiz.ReadProblemsFile()
	problems := quiz.ParseProblems(records)
	totalQuestions, correctAnswers := quiz.StartQuiz(problems)
	fmt.Println(totalQuestions, correctAnswers)

}
