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
	if score1 >= 9 {
		promoters = promoters + 1
	}
	if score1 <= 6 {
		detractors = detractors + 1
	}
	if score2 >= 9 {
		promoters = promoters + 1
	}
	if score2 <= 6 {
		detractors = detractors + 1
	}
	if score3 >= 9 {
		promoters = promoters + 1
	}
	if score3 <= 6 {
		detractors = detractors + 1
	}
	nps := (promoters - detractors) * 100 / 3
	fmt.Println(nps)
	//if - условия
	//boolean - тип данных
}
