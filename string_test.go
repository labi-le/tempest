package tempest

import (
	"testing"
)

func TestReplaceString(t *testing.T) {
	type args struct {
		s            string
		replacements []Replacement
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestReplaceString",
			args: args{
				s: "hello {name}",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "world",
					},
				},
			},
			want: "hello world",
		},

		{
			name: "TestReplaceString multiple",
			args: args{
				s: "hello {name}, {name2}",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "world",
					},

					{
						Tag:   "name2",
						Value: "universe",
					},
				},
			},
			want: "hello world, universe",
		},

		{
			name: "TestReplaceString empty string",
			args: args{
				s:            "hello {name}",
				replacements: []Replacement{},
			},
			want: "hello {name}",
		},

		{
			name: "TestReplaceString overlapping tags",
			args: args{
				s: "hello {name}, {name2}",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "world",
					},
				},
			},
			want: "hello world, {name2}",
		},

		{
			name: "TestReplaceString recursive replacement",
			args: args{
				s: "hello {name}",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "{greeting}",
					},
					{
						Tag:   "greeting",
						Value: "{tree}",
					},
					{
						Tag:   "tree",
						Value: "universe",
					},
				},
			},
			want: "hello universe",
		},
		{
			name: "TestReplaceString case sensitivity",
			args: args{
				s: "hello {Name}",
				replacements: []Replacement{
					{
						Tag:   "Name",
						Value: "world",
					},
				},
			},
			want: "hello world",
		},
		{
			name: "TestReplaceString string is a tag",
			args: args{
				s: "{name}",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "world",
					},
				},
			},
			want: "world",
		},
		{
			name: "TestReplaceString nested tags",
			args: args{
				s: "{hello {name}}",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "world",
					},
				},
			},
			want: "{hello world}",
		},

		{
			name: "TestReplaceString non-existing tag",
			args: args{
				s: "hello {non_existing}",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "world",
					},
				},
			},
			want: "hello {non_existing}",
		},
		{
			name: "TestReplaceString malformed tag",
			args: args{
				s: "hello {name",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "world",
					},
				},
			},
			want: "hello {name",
		},
		{
			name: "TestReplaceString replace with empty string",
			args: args{
				s: "hello {name}",
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "",
					},
				},
			},
			want: "hello ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceString(tt.args.s, tt.args.replacements); got != tt.want {
				t.Errorf("ReplaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
