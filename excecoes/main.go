package main

import (
	"errors"
	"fmt"
	"math"
)

// NegativeSqrtError implementa a interface de exceção da linguagem Go
//
//	type error interface {
//	    Error() string
//	}
type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, NegativeSqrtError(f)
	}
	return math.Sqrt(f), nil
}

func main() {
	if _, err := Sqrt(-90); err != nil {
		var negError NegativeSqrtError
		if errors.As(err, &negError) {
			fmt.Println(err)
		} else {
			fmt.Println("unexpected error")
		}
		return
	}
}
