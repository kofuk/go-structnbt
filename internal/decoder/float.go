package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"
	"unsafe"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagFloatDecoder struct{}

var _ TypedDecoder = (*TagFloatDecoder)(nil)

func (d *TagFloatDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	if v != nil {
		if !v.CanFloat() {
			return &errors.TypeMismatchError{
				TagType: types.TagFloat,
				DstType: v.Type(),
			}
		}
		tmp := binary.BigEndian.Uint32(buf)
		binary.LittleEndian.PutUint32(buf, tmp)
		v.SetFloat(float64(*(*float32)(unsafe.Pointer(&buf[0]))))
	}
	return nil
}
