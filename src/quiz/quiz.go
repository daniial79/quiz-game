package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/daniial79/quiz-game/src/problem"
	"github.com/daniial79/quiz-game/src/utils"
	"os"
)

func Welcome() {
	fmt.Println("Welcome to this quiz program")
	fmt.Println("------------------------")
}

func ReadProblemsFile() [][]string {
	f, err := os.Open("./src/problems.csv")
	if err != nil {
		utils.LogErr("Unable to open problems file", err)
	}

	defer func(*os.File) {
		err := f.Close()

		if err != nil {
			utils.LogErr("Unable to close problems file", err)
		}
	}(f)

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		utils.LogErr("Unable to parse problems to string", err)
	}

	return records
}

func ParseProblems(records [][]string) []problem.Problem {
	problems := make([]problem.Problem, len(records))

	for i := range records {
		problems[i] = problem.NewProblem(
			records[i][0],
			records[i][1],
		)
	}

	return problems
}

func StartQuiz(problems []problem.Problem) (problemCount int, correctAns int) {
	var correctAnswers int

	scanner := bufio.NewScanner(os.Stdin)
	for n, p := range problems {

		var fetchedAns string
		fmt.Printf("%d) %s ? ", n+1, p.GetQuestion())
		if scanner.Scan() {
			fetchedAns = scanner.Text()
		}

		if fetchedAns == p.GetAnswer() {
			correctAnswers++
		}
	}

	return len(problems), correctAnswers
}

func PrintResult(totalQuestions, correctAnswers int) {
	fmt.Println("------------------------")
	fmt.Printf("You have answered %d quesionts correctly out of %d\n", correctAnswers, totalQuestions)
}
