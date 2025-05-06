package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagByteArrayDecoder struct{}

var _ TypedDecoder = (*TagByteArrayDecoder)(nil)

func (d *TagByteArrayDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	lenBuf := make([]byte, 4)
	if _, err := io.ReadFull(r, lenBuf); err != nil {
		return err
	}
	arrayLen := int(binary.BigEndian.Uint32(lenBuf))

	var result []byte
	for i := 0; i < arrayLen; i++ {
		value, err := r.ReadByte()
		if err != nil {
			return err
		}
		result = append(result, value)
	}
	if v != nil {
		if v.Type() != reflect.SliceOf(reflect.TypeOf(byte(0))) {
			return &errors.TypeMismatchError{
				TagType: types.TagByteArray,
				DstType: v.Type(),
			}
		}
		v.Set(reflect.ValueOf(result))
	}
	return nil
}
