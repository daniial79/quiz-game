package problem

type Problem struct {
	question string
	answer   string
}

func NewProblem(q, a string) Problem {
	return Problem{
		question: q,
		answer:   a,
	}
}

func (p Problem) GetQuestion() string {
	return p.question
}

func (p Problem) GetAnswer() string {
	return p.answer
}
