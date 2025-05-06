package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagIntArrayDecoder struct{}

var _ TypedDecoder = (*TagIntArrayDecoder)(nil)

func (d *TagIntArrayDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	lenBuf := make([]byte, 4)
	if _, err := io.ReadFull(r, lenBuf); err != nil {
		return err
	}
	arrayLen := int(binary.BigEndian.Uint32(lenBuf))

	var result []int32
	buf := make([]byte, 4)
	for i := 0; i < arrayLen; i++ {
		if _, err := io.ReadFull(r, buf); err != nil {
			return err
		}
		result = append(result, int32(binary.BigEndian.Uint32(buf)))
	}
	if v != nil {
		if v.Type() != reflect.SliceOf(reflect.TypeOf(int32(0))) {
			return &errors.TypeMismatchError{
				TagType: types.TagIntArray,
				DstType: v.Type(),
			}
		}
		v.Set(reflect.ValueOf(result))
	}
	return nil
}
