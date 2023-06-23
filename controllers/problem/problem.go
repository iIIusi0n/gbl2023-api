package problem

import (
	"gbl-api/data"
)

func GetBoothProblems(bid string) ([]Problem, error) {
	var problems []Problem
	db := data.GetDatabase()

	err := db.Where("bid = ?", bid).Find(&problems).Error

	return problems, err
}

func CheckAnswer(pid, answer string) int {
	var problem Problem
	db := data.GetDatabase()

	db.Where("pid = ?", pid).First(&problem)

	if problem.Answer == answer {
		return problem.Score
	}

	return 0
}
