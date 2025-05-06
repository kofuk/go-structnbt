package errors

import (
	"fmt"
	"reflect"

	"github.com/kofuk/go-structnbt/types"
)

type TypeMismatchError struct {
	TagType types.TagType
	DstType reflect.Type
}

func (e *TypeMismatchError) Error() string {
	return fmt.Sprintf(
		"value of type %s can't hold %s",
		e.DstType.String(),
		e.TagType.String(),
	)
}

type DepthLimitError struct {
	Level int
	Limit int
}

func (e *DepthLimitError) Error() string {
	return fmt.Sprintf("depth limit reached: %d (limit: %d)", e.Level, e.Limit)
}
