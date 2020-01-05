package main

func main() {
	// problem 1: как хранить много данных
	// go -> массив заполняется нулевыми значениями (по умолчанию)
	// neutrals? - practice vs theory
	// ctrl + alt + v с виделением - позволяет создать локальную ..
	// shift + f6
	// problem x: auto-testing
	// ctrl + alt + shift + левый курсор мыши
	// nps = 100
	// _ - don't care (все равно)

	// if - условия
	// boolean - тип данных
	// refactoring: улучшение структуры код без модификации поведение
	//slice - динамически изменяемый "список"
	//n/ps -> ctrl + shift + T
}
func nps(scores []int) int {
	promoters := 0
	detractors := 0
	promotersLowerBound := 9
	detractorsLowerBound := 6

	for _, value := range scores{
		if value >= promotersLowerBound {
			promoters = promoters + 1
		}
		if value <= detractorsLowerBound {
			detractors = detractors + 1
		}
	}
	nps := (promoters - detractors) * 100 / len(scores)
	return nps
}