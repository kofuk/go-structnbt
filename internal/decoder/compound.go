package decoder

import (
	"bufio"
	"reflect"

	"github.com/kofuk/go-structnbt/errors"
	"github.com/kofuk/go-structnbt/types"
)

type TagCompoundDecoder struct {
	levelLimit int
}

var _ TypedDecoder = (*TagCompoundDecoder)(nil)

func (d *TagCompoundDecoder) Decode(r *bufio.Reader, v *reflect.Value, level int) error {
	if d.levelLimit > 0 && level >= d.levelLimit {
		return &errors.DepthLimitError{
			Level: level,
			Limit: d.levelLimit,
		}
	}

	if v != nil && v.Kind() != reflect.Struct {
		return &errors.TypeMismatchError{
			TagType: types.TagCompound,
			DstType: v.Type(),
		}
	}

	fields := make(map[string]StructField)
	if v != nil {
		fields = GetStructFields(v.Type())
	}

	for {
		ty, err := readTagType(r)
		if err != nil {
			return err
		}

		if ty == types.TagEnd {
			break
		}

		name, err := readName(r)
		if err != nil {
			return err
		}

		typeReader, err := getTypeReader(ty, d.levelLimit)
		if err != nil {
			return err
		}

		var value *reflect.Value
		if field, ok := fields[name]; ok {
			tmp := v.Field(field.Index)
			value = &tmp
		}

		if err := typeReader.Decode(r, value, level+1); err != nil {
			return err
		}
	}

	return nil
}
