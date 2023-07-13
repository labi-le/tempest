package tempest

import (
	"reflect"
	"strings"
)

const (
	tag      = "template"
	maxDepth = 10
)

// Replacement is a struct that holds a tag and a value
// Tag which will be searched in the data structure
// Value that will replace the tag
type Replacement struct {
	Tag   string
	Value string
}

type Config struct {
	// OpenTag property is the opening tag e.g. "{"
	OpenTag string
	// CloseTag property is the closing tag e.g. "}"
	CloseTag string
}

// config is the default configuration
var config = Config{
	OpenTag:  "{",
	CloseTag: "}",
}

// SetConfig sets the global configuration
func SetConfig(c Config) {
	config = c
}

// replace replaces all occurrences of tag in src with val
func replace(src string, tag string, val string) string {
	return strings.ReplaceAll(src, config.OpenTag+tag+config.CloseTag, val)
}

// replace replaces all occurrences of tag in src with val
func replaceWithDepth(src string, tag string, val string, depth int) string {
	if depth > maxDepth {
		panic("max depth reached")
	}
	return replace(src, tag, val)
}

// isPtr checks if val is a pointer
func isPtr(val reflect.Value) bool {
	return val.Kind() == reflect.Ptr
}

// isStruct checks if val is a struct
func isStruct(val reflect.Value) bool {
	return val.Elem().Kind() == reflect.Struct
}
