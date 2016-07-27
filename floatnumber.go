package parse

import "fmt"

// Make sure that when we marshal there will always be a precise point after the number (e.g. - 1.000)
type FloatNumber float64

func (n FloatNumber) MarshalJSON() ([]byte, error) {

	return []byte(fmt.Sprintf("%f", n)), nil
}
