package arrays_n_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	response := Sum(numbers)

	expected_response := 15

	if response != expected_response {
		t.Errorf("response '%d', expected '%d' from '%v", response, expected_response, numbers)
	}
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5}

	response := SumAll(numbers1, numbers2)

	expected_response := []int{6, 9}

	if !reflect.DeepEqual(response, expected_response) {
		t.Errorf("response '%v', expected '%v'", response, expected_response)
	}
}

func TestSumAllTheRest(t *testing.T) {
	verifySums := func(t *testing.T, response, expected_response []int) {
		t.Helper()

		if !reflect.DeepEqual(response, expected_response) {
			t.Errorf("response '%v', expected '%v'", response, expected_response)

		}
	}

	t.Run("when a valid slice is given, sum all the rest", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{4, 5}

		response := SumAllTheRest(numbers1, numbers2)

		expected_response := []int{5, 5}

		verifySums(t, response, expected_response)
	})

	t.Run("when a empty slice is given, return 0 for it", func(t *testing.T) {
		numbers1 := []int{}
		numbers2 := []int{4, 5}

		response := SumAllTheRest(numbers1, numbers2)

		expected_response := []int{0, 5}

		verifySums(t, response, expected_response)
	})
}
