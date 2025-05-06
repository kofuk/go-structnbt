package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagLongDecoder struct{}

var _ TypedDecoder = (*TagLongDecoder)(nil)

func (d *TagLongDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	if v != nil {
		if !v.CanInt() {
			return &errors.TypeMismatchError{
				TagType: types.TagLong,
				DstType: v.Type(),
			}
		}
		value := int64(binary.BigEndian.Uint64(buf))
		v.SetInt(value)
	}
	return nil
}
