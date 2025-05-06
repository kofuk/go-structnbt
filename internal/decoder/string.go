package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagStringDecoder struct{}

var _ TypedDecoder = (*TagStringDecoder)(nil)

func (d *TagStringDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	lenBuf := make([]byte, 2)
	if _, err := io.ReadFull(r, lenBuf); err != nil {
		return err
	}
	valueLen := binary.BigEndian.Uint16(lenBuf)

	value := make([]byte, int(valueLen))
	if _, err := io.ReadFull(r, value); err != nil {
		return err
	}
	if v != nil {
		if v.Kind() != reflect.String {
			return &errors.TypeMismatchError{
				TagType: types.TagString,
				DstType: v.Type(),
			}
		}
		v.SetString(string(value))
	}
	return nil
}
