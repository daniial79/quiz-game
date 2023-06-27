package app

import (
	"github.com/daniial79/quiz-game/src/quiz"
)

func Start() {
	quiz.Welcome()
	records := quiz.ReadProblemsFile()
	problems := quiz.ParseProblems(records)
	totalQuestions, correctAnswers := quiz.StartQuiz(problems)
	quiz.PrintResult(totalQuestions, correctAnswers)
}
