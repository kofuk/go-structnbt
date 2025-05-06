package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagShortDecoder struct{}

var _ TypedDecoder = (*TagShortDecoder)(nil)

func (d *TagShortDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	if v != nil {
		if !v.CanInt() {
			return &errors.TypeMismatchError{
				TagType: types.TagShort,
				DstType: v.Type(),
			}
		}
		value := int64(int16(binary.BigEndian.Uint16(buf)))
		v.SetInt(value)
	}
	return nil
}
