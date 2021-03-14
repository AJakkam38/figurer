package calc

import "testing"

type inputTest struct {
	arg1, arg2 string
}

var inputs = []inputTest{
	inputTest{"2", "3"},
	inputTest{"-4", "8"},
	inputTest{"-2", "-2"},
	inputTest{"2.5", "6"},
}

func TestAdd(t *testing.T) {
	inputResults := []float64{5, 4, -4, 8.5}
	for i, test := range inputs {
		if output := Calculate(test.arg1, test.arg2, "+"); output != inputResults[i] {
			t.Errorf("Output %f not equal to expected %f", output, inputResults[i])
		}
	}
}

func TestSubtract(t *testing.T) {
	inputResults := []float64{-1, -12, 0, -3.5}
	for i, test := range inputs {
		if output := Calculate(test.arg1, test.arg2, "-"); output != inputResults[i] {
			t.Errorf("Output %f not equal to expected %f", output, inputResults[i])
		}
	}
}

func TestMultiply(t *testing.T) {
	inputResults := []float64{6, -32, 4, 15}
	for i, test := range inputs {
		if output := Calculate(test.arg1, test.arg2, "*"); output != inputResults[i] {
			t.Errorf("Output %f not equal to expected %f", output, inputResults[i])
		}
	}
}

func TestDivide(t *testing.T) {
	inputResults := []float64{0.67, -0.5, 1, 0.42}
	for i, test := range inputs {
		if output := Calculate(test.arg1, test.arg2, "/"); output != inputResults[i] {
			t.Errorf("Output %f not equal to expected %f", output, inputResults[i])
		}
	}
}
