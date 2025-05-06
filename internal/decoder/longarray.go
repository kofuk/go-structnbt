package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagLongArrayDecoder struct{}

var _ TypedDecoder = (*TagLongArrayDecoder)(nil)

func (d *TagLongArrayDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	lenBuf := make([]byte, 4)
	if _, err := io.ReadFull(r, lenBuf); err != nil {
		return err
	}
	arrayLen := int(binary.BigEndian.Uint32(lenBuf))

	var result []int64
	buf := make([]byte, 8)
	for i := 0; i < arrayLen; i++ {
		if _, err := io.ReadFull(r, buf); err != nil {
			return err
		}
		result = append(result, int64(binary.BigEndian.Uint64(buf)))
	}
	if v != nil {
		if v.Type() != reflect.SliceOf(reflect.TypeOf(int64(0))) {
			return &errors.TypeMismatchError{
				TagType: types.TagLongArray,
				DstType: v.Type(),
			}
		}
		v.Set(reflect.ValueOf(result))
	}
	return nil
}
