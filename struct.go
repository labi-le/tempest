package tempest

import (
	"reflect"
	"strings"
)

func ReplaceStructByTag(v any, replacements []Replacement) {
	val := reflect.ValueOf(v)

	if !isPtr(val) || !isStruct(val) {
		panic("v must be a pointer to a struct")
	}

	val = val.Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := typ.Field(i)

		if valueField.Kind() == reflect.String && valueField.CanSet() {
			oldValue := ""
			newValue := valueField.String()

			depth := 0
			for oldValue != newValue {
				oldValue = newValue

				for _, name := range strings.Split(typeField.Tag.Get(tag), ",") {
					for _, replacement := range replacements {
						if replacement.Tag == name {
							newValue = replaceWithDepth(newValue, replacement.Tag, replacement.Value, depth)
							depth++
						}
					}
				}
			}

			valueField.SetString(newValue)
		}
	}
}

func ReplaceStruct(v any, replacements []Replacement) {
	val := reflect.ValueOf(v)

	if !isPtr(val) || !isStruct(val) {
		panic("v must be a pointer to a struct")
	}

	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)

		if valueField.Kind() == reflect.String && valueField.CanSet() {
			oldValue := ""
			newValue := valueField.String()

			depth := 0
			for oldValue != newValue {
				oldValue = newValue

				for _, replacement := range replacements {
					newValue = replaceWithDepth(newValue, replacement.Tag, replacement.Value, depth)
					depth++
				}
			}

			valueField.SetString(newValue)
		}
	}
}
