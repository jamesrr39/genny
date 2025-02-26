package math

import "github.com/jamesrr39/genny/generic"

type NumberType generic.Number

// NumberTypeMax gets the maximum number from the
// two specified.
func NumberTypeMax(a, b NumberType) NumberType {
	if a > b {
		return a
	}
	return b
}
