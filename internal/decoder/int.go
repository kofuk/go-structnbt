package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagIntDecoder struct{}

var _ TypedDecoder = (*TagIntDecoder)(nil)

func (d *TagIntDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	if v != nil {
		if !v.CanInt() {
			return &errors.TypeMismatchError{
				TagType: types.TagInt,
				DstType: v.Type(),
			}
		}
		value := int64(int32(binary.BigEndian.Uint32(buf)))
		v.SetInt(value)
	}
	return nil
}
