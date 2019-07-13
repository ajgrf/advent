package euler

import (
	"fmt"
	"reflect"
	"testing"
)

var answers = []struct {
	problem interface{}
	answer  string
}{
	1:  {Problem001, "233168"},
	2:  {Problem002, "4613732"},
	3:  {Problem003, "6857"},
	4:  {Problem004, "906609"},
	5:  {Problem005, "232792560"},
	6:  {Problem006, "25164150"},
	7:  {Problem007, "104743"},
	8:  {Problem008, "40824"},
	9:  {Problem009, "31875000"},
	10: {Problem010, "142913828922"},
	11: {Problem011, "70600674"},
	12: {Problem012, "76576500"},
	13: {Problem013, "5537376230"},
	14: {Problem014, "837799"},
	16: {Problem016, "1366"},
	17: {Problem017, "21124"},
	19: {Problem019, "171"},
	20: {Problem020, "648"},
	24: {Problem024, "2783915460"},
	25: {Problem025, "4782"},
	29: {Problem029, "9183"},
	36: {Problem036, "872187"},
	48: {Problem048, "9110846700"},
}

func TestProblems(t *testing.T) {
	_ = t
	for i := range answers {
		problem := answers[i].problem
		answer := answers[i].answer

		v := reflect.ValueOf(problem)
		if v.IsValid() && v.Kind() == reflect.Func &&
			v.Type().NumIn() == 0 && v.Type().NumOut() == 1 {

			result := fmt.Sprint(v.Call(nil)[0].Interface())
			if answer == "" {
				fmt.Printf("%3d. %s\n", i, result)
			} else if result != answer {
				t.Errorf("%3d. expected %s, got %s", i, answer, result)
			}

		}
	}
}
