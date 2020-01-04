package main

import "fmt"

func main() {
	// problem 1: как хранить много данных
	score1 := 0
	score2 := 7
	score3 := 0

	// nps = 100
	promoters := 0
	detractors := 0
	//neutrals? - practice vs theory
	//ctrl + alt + v с виделением - позволяет создать локальную ..
	//shift + f6
	// problem x: auto-testing
	//ctrl + alt + shift + левый курсор мыши
	promotersLowerBound := 9
	if score1 >= promotersLowerBound {
		promoters = promoters + 1
	}
	detractorsLowerBound := 6
	if score1 <= detractorsLowerBound {
		detractors = detractors + 1
	}
	if score2 >= promotersLowerBound {
		promoters = promoters + 1
	}
	if score2 <= detractorsLowerBound {
		detractors = detractors + 1
	}
	if score3 >= promotersLowerBound {
		promoters = promoters + 1
	}
	if score3 <= detractorsLowerBound {
		detractors = detractors + 1
	}
	nps := (promoters - detractors) * 100 / 3
	fmt.Println(nps)
	//if - условия
	//boolean - тип данных
	//refactoring: улучшение структуры код без модификации поведение
}
