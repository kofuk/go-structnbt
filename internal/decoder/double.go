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

type TagDoubleDecoder struct{}

var _ TypedDecoder = (*TagDoubleDecoder)(nil)

func (d *TagDoubleDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	if v != nil {
		if !v.CanFloat() {
			return &errors.TypeMismatchError{
				TagType: types.TagDouble,
				DstType: v.Type(),
			}
		}
		tmp := binary.BigEndian.Uint64(buf)
		binary.LittleEndian.PutUint64(buf, tmp)
		v.SetFloat(*(*float64)(unsafe.Pointer(&buf[0])))
	}
	return nil
}
