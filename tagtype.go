package structnbt

//go:generate go run golang.org/x/tools/cmd/stringer@v0.32.0 -type=TagType -output=tagtype_string.go

type TagType byte

const (
	TagEnd TagType = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
	TagLongArray
)
