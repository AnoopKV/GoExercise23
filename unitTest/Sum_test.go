package uTest

import (
	"fmt"
	"reflect"
	"testing"
)

type test[T int | string | float64] struct {
	val1 T
	val2 T
	want T
}

func execute[T int | string | float64](tests map[string]test[T], t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Addition(tc.val1, tc.val2)
			//fmt.Println(tc.want, got)
			if !reflect.DeepEqual(tc.want, got) {
				fmt.Println(t.Failed()) //Q4: Why is it saying false? it supposed to be true!
				//t.Skipf("Skipping testcase, Failed testcase details :: "+"expected: %#v, got: %#v", tc.want, got)
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

func TestAddition(t *testing.T) {
	/*
		//Q1: How can I create anonymous struct with generic type? []struct[T]{...}{...} throws error
		//Q2: How can I combine all testcases under one tests, as given below _tests?
		//Q3: How can I create above method execute(...) within TestAddition, I tried with execute:=func [T  int | string | float64](), thrown error
		_tests := map[string]test[T]{
			"Test1": {val1: "1", val2: "2", want: "12"},
			"Test2": {val1: "1", val2: "3", want: "13"},
			"Test3": {val1: 1, val2: 2, want: 3},
			"Test4": {val1: 2, val2: 3, want: 5},
			"Test5": {val1: 1.0, val2: 2.0, want: 3.0},
			"Test6": {val1: 2.0, val2: 3.0, want: 5.0},
		}
		execute(_tests, t)
	*/

	tests := map[string]test[string]{
		"Test-1": {val1: "1", val2: "2", want: "112"},
		"Test-2": {val1: "1", val2: "3", want: "13"},
	}
	execute(tests, t)

	test1 := map[string]test[int]{
		"Test-3": {val1: 1, val2: 2, want: 3},
		"Test-4": {val1: 2, val2: 3, want: 5},
	}
	execute(test1, t)

	test3 := map[string]test[float64]{
		"Test-5": {val1: 1.0, val2: 2.0, want: 3.0},
		"Test-6": {val1: 2.0, val2: 3.0, want: 5.0},
	}
	execute(test3, t)
}
