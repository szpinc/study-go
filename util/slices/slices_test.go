package slices

import "testing"

func TestContains(t *testing.T) {

	s := []string{"a", "b", "c", "d"}

	resultTrue := Contains(s, "a")

	resultFalse := Contains(s, "asd")

	println("resultTrue:", resultTrue)

	println("resultFalse:", resultFalse)

}
