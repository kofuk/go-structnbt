package structnbt

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"reflect"

	"github.com/kofuk/go-structnbt/internal/decoder"
)

type Decoder struct {
	r     *bufio.Reader
	limit int
}

type DecoderOption func(*Decoder)

func WithMaxDepth(limit int) DecoderOption {
	return func(d *Decoder) {
		d.limit = limit
	}
}

func NewDecoder(r io.Reader, options ...DecoderOption) *Decoder {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}

	decoder := &Decoder{
		r: br,
	}

	for _, opt := range options {
		opt(decoder)
	}

	return decoder
}

func (dec *Decoder) Decode(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errors.New("v must be a pointer and not be nil")
	}

	decoder, err := decoder.GetInitialDecoder(dec.r, dec.limit)
	if err != nil {
		return err
	}

	valueElem := rv.Elem()

	return decoder.Decode(dec.r, &valueElem, 0)
}

func Unmarshal(data []byte, v any) error {
	return NewDecoder(bytes.NewBuffer(data)).Decode(v)
}
