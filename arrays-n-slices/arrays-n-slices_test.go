package arrays_n_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("colection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		response := Sum(numbers)

		expected_response := 15

		if response != expected_response {
			t.Errorf("response '%d', expected '%d' from '%v", response, expected_response, numbers)
		}
	})
}
