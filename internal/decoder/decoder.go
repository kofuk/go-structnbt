package decoder

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/types"
)

type TypedDecoder interface {
	Decode(r *bufio.Reader, v *reflect.Value, level int) error
}

func readTagType(r *bufio.Reader) (types.TagType, error) {
	ty, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	return types.TagType(ty), nil
}

func getTypeReader(ty types.TagType, levelLimit int) (TypedDecoder, error) {
	switch ty {
	case types.TagEnd:
		return nil, errors.New("invalid TAG_End")
	case types.TagByte:
		return &TagByteDecoder{}, nil
	case types.TagShort:
		return &TagShortDecoder{}, nil
	case types.TagInt:
		return &TagIntDecoder{}, nil
	case types.TagLong:
		return &TagLongDecoder{}, nil
	case types.TagFloat:
		return &TagFloatDecoder{}, nil
	case types.TagDouble:
		return &TagDoubleDecoder{}, nil
	case types.TagByteArray:
		return &TagByteArrayDecoder{}, nil
	case types.TagString:
		return &TagStringDecoder{}, nil
	case types.TagList:
		return &TagListDecoder{
			levelLimit: levelLimit,
		}, nil
	case types.TagCompound:
		return &TagCompoundDecoder{
			levelLimit: levelLimit,
		}, nil
	case types.TagIntArray:
		return &TagIntArrayDecoder{}, nil
	case types.TagLongArray:
		return &TagLongArrayDecoder{}, nil
	default:
		return nil, fmt.Errorf("invalid tag type: %d", ty)
	}
}

func readName(r *bufio.Reader) (string, error) {
	nameLenBuf := make([]byte, 2)
	if _, err := io.ReadFull(r, nameLenBuf); err != nil {
		return "", err
	}
	nameLen := binary.BigEndian.Uint16(nameLenBuf)

	name := make([]byte, int(nameLen))
	if _, err := io.ReadFull(r, name); err != nil {
		return "", err
	}
	return string(name), nil
}

func GetInitialDecoder(r *bufio.Reader, levelLimit int) (TypedDecoder, error) {
	nbtTy, err := readTagType(r)
	if err != nil {
		return nil, err
	}

	// Dispose tag name of top tag
	readName(r)

	return getTypeReader(nbtTy, levelLimit)
}
