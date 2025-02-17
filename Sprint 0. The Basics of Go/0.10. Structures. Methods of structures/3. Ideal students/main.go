package main

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (e Student) IsExcellentStudent() bool {
	if float64(e.solvedProblems)*e.scoreForOneTask >= e.passingScore {
		return true
	}
	return false
}
