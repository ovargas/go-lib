package constant

import "golang.org/x/exp/constraints"

type (
	Constant interface {
		constraints.Complex | constraints.Ordered | ~bool
	}
)

func AsPointer[T Constant](value T) *T {
	return &value
}
