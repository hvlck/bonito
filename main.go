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

// Separates the string at every three significant places with sep.
func (b *Bonito[T]) AsSeparated(sep string) *Bonito[T] {
	if len(b.str) > 3 {
		leading_places := len(b.str) % 3
		new := ""

		if leading_places != 0 {
			new += b.str[0:leading_places] + sep
		}

		for i := leading_places; i < len(b.str); i += 3 {
			new += b.str[i : i+3]
			if i+3 < len(b.str) {
				new += sep
			}
		}

		b.str = new
	}

	return b
}


func (b *Bonito[T]) Trunc(prec int) *Bonito[T] {
	if prec >= len(b.str) {
		return b
	}

	if prec+1 > len(b.str) {
		prec--
	}
	b.significand = len(b.str) - prec

	conv, err := strconv.ParseFloat(b.str[0:prec+1], 64)
	conv = conv / 10
	b.str = fmt.Sprint(math.Round(conv))

	if err != nil {
		b.str = b.str[0 : prec+1]
	}

	return b
}
