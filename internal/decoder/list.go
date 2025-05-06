package decoder

import (
	"bufio"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
)

type TagListDecoder struct {
	levelLimit int
}

func (d *TagListDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	ty, err := readTagType(r)
	if err != nil {
		return err
	}

	lenBuf := make([]byte, 4)
	if _, err := io.ReadFull(r, lenBuf); err != nil {
		return err
	}
	listLen := int(binary.BigEndian.Uint32(lenBuf))

	if listLen == 0 {
		// Minecraft can generate lists of TAG_End with length 0.
		// We don't care about elements' type if it's an empty list.
		return nil
	}

	typeDecoder, err := getTypeReader(ty, d.levelLimit)
	if err != nil {
		return err
	}

	if v == nil {
		for i := 0; i < listLen; i++ {
			if err := typeDecoder.Decode(r, nil, level); err != nil {
				return err
			}
		}
	} else {
		if v.Kind() != reflect.Slice {
			return &errors.TypeMismatchError{
				TagType: ty,
				DstType: v.Type(),
			}
		}
		elemType := v.Type().Elem()
		slice := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0)
		for i := 0; i < listLen; i++ {
			v := reflect.New(elemType).Elem()
			if err := typeDecoder.Decode(r, &v, level); err != nil {
				return err
			}
			slice = reflect.Append(slice, v)
		}
		v.Set(slice)
	}

	return nil
}
