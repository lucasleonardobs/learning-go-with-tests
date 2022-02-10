package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)

	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', receive '%d'", expected, sum)
	}
}

func ExampleAdd() {
	soma := Add(1, 5)
	fmt.Println(soma)
	// Output: 6
}
