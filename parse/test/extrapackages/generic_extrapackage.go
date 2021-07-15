package extrapackages

import "github.com/jamesrr39/genny/generic"

type ForeignType generic.Type

func ForeignTypeSayHello(a ForeignType) ForeignType {
	return a
}
