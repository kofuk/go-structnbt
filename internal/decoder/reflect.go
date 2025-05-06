package decoder

import "reflect"

type StructField struct {
	Name  string
	Index int
}

func GetStructFields(ty reflect.Type) map[string]StructField {
	fields := make(map[string]StructField)

	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		name := field.Tag.Get("nbt")
		if name == "" {
			name = field.Name
		}

		fields[name] = StructField{
			Name:  name,
			Index: i,
		}
	}

	return fields
}
