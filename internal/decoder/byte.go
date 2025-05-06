package decoder

import (
	"bufio"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagByteDecoder struct{}

var _ TypedDecoder = (*TagByteDecoder)(nil)

func (d *TagByteDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	value, err := r.ReadByte()
	if err != nil {
		return err
	}
	if v != nil {
		if !v.CanInt() {
			return &errors.TypeMismatchError{
				TagType: types.TagByte,
				DstType: v.Type(),
			}
		}
		v.SetInt(int64(int8(value)))
	}
	return nil
}
