package unexported

import (
	"fmt"

	"github.com/jamesrr39/genny/generic"
)

type secret generic.Type

func secretInspect(s secret) string {
	return fmt.Sprintf("%#v", s)
}
