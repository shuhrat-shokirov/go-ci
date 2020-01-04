package main

import (
	"flag"
	"fmt"
)

func main() {
	// problem 1: как хранить много данных
	//go -> массив заполняется нулевыми значениями (по умолчанию)
	scores := [3]int{10, 7, 10}

	// nps = 100
	promoters := 0
	detractors := 0
	//neutrals? - practice vs theory
	//ctrl + alt + v с виделением - позволяет создать локальную ..
	//shift + f6
	// problem x: auto-testing
	//ctrl + alt + shift + левый курсор мыши
	promotersLowerBound := 9
	detractorsLowerBound := 6
	\\_ - don't care (все равно)
	for _, value := range scores{
		if value >= promotersLowerBound {
			promoters = promoters + 1
		}
		if value <= detractorsLowerBound {
			detractors = detractors + 1
		}
	}
	nps := (promoters - detractors) * 100 / len(scores)
	fmt.Println(nps)
	//if - условия
	//boolean - тип данных
	//refactoring: улучшение структуры код без модификации поведение
}
