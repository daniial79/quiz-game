package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/daniial79/quiz-game/src/utils"
	"os"
)

type Problem struct {
	Question string
	Answer   string
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

func ParseProblems(records [][]string) []Problem {
	problems := make([]Problem, len(records))

	for i := range records {
		problems[i] = Problem{
			Question: records[i][0],
			Answer:   records[i][1],
		}
	}

	return problems
}

func StartQuiz(problems []Problem) (problemCount int, correctAns int) {
	var correctAnswers int

	scanner := bufio.NewScanner(os.Stdin)
	for n, p := range problems {

		var fetchedAns string
		fmt.Printf("%d) %s ? ", n+1, p.Question)
		if scanner.Scan() {
			fetchedAns = scanner.Text()
		}

		if fetchedAns == p.Answer {
			correctAnswers++
		}
	}

	return len(problems), correctAnswers
}
