package main

import "testing"
//	debugger и запуск отсюда тоже работают
func Test_nps(t *testing.T) {
	/*scores := []int{10, 7, 10, 10, 10}
	// want - ожидаемый результат
	// got - тоб что реально получили
	want := 100
	got := nps(scores)
	if want !=got {
		t.Error("nps with args:", scores, "want:", want, "got:", got)
	}
*/
	// testCases - тестовые сценарии
	testCases := []struct{
		name string
		score []int
		want int
	}{
		{"All promoters", []int{10, 10, 10}, 100},
		{"One promoter", []int{10}, 100},
		{"All detractors", []int{0, 0, 0},-100},
		{"All neutrals",[]int{8, 8, 8}, 0},
		{"Mixed",[]int{10, 10, 8}, 66},
	}
	for _, testCase := range testCases {
		got := nps(testCase.score)
		if got != testCase.want{
			t.Error("nps with test:", testCase.name, "got:", got, "want:", testCase.want)
		}
	}
}