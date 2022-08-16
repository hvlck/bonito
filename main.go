package bonito

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		big.Float | big.Int | big.Rat
}

type Bonito[T Number] struct {
	number *T
	str    string
	// Significand of the number, if rounded. Only set in calls to Bonito.WithPrecision()
	significand int
}

func Pretty[T Number](number T) *Bonito[T] {
	return &Bonito[T]{
		number: &number,
		str:    fmt.Sprint(number),
	}
}
