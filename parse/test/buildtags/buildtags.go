// +build x,y z
// +build genny

package buildtags

import (
	"fmt"

	"github.com/jamesrr39/genny/generic"
)

type _t_ generic.Type

func _t_Print(t _t_) {
	fmt.Println(t)
}
