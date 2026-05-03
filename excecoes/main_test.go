package main

import (
	"errors"
	"testing"
)

var negError NegativeSqrtError

func TestNegativeInput(t *testing.T) {
	f := float64(-90)
	_, err := Sqrt(f)
	if !errors.As(err, &negError) {
		t.Error(err)
	}
}