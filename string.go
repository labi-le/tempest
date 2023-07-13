package tempest

// ReplaceString replaces all tags in a string with the given replacements
//
// e.g:
// "hello {name}" with {name: "world"} will return "hello world"
//
// "hello {name}, {name2}" with {name: "world", name2: "universe"} will return "hello world, universe"
func ReplaceString(s string, replacements []Replacement) string {
	for _, v := range replacements {
		s = replace(s, v.Tag, v.Value)
	}
	return s
}
