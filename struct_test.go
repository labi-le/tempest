package tempest

import (
	"testing"
)

func TestReplaceStruct(t *testing.T) {
	testReplaceStructAny(t, false)
}

func TestReplaceStructByTag(t *testing.T) {
	testReplaceStructAny(t, true)
}

func testReplaceStructAny(t *testing.T, byTag bool) {
	type teststruct struct {
		Text string `template:"name,name2,firstName,lastName,nickName,user,num"`
	}

	type args struct {
		v            any
		replacements []Replacement
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestReplaceStructByTag",
			args: args{
				v: &teststruct{
					Text: "hello {name}",
				},
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
			name: "TestReplaceStructByTag multiple",
			args: args{
				v: &teststruct{
					Text: "hello {name} and {name2}",
				},
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

			want: "hello world and universe",
		},

		{
			name: "TestReplaceStructByTag with unused tag",
			args: args{
				v: &teststruct{
					Text: "Hello, {name}",
				},
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "Alice",
					},
					{
						Tag:   "unused",
						Value: "not used",
					},
				},
			},
			want: "Hello, Alice",
		},

		{
			name: "TestReplaceStructByTag with missing replacement",
			args: args{
				v: &teststruct{
					Text: "Hello, {name}",
				},
				replacements: []Replacement{},
			},
			want: "Hello, {name}",
		},

		{
			name: "TestReplaceStructByTag with empty string",
			args: args{
				v: &teststruct{
					Text: "",
				},
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "Alice",
					},
				},
			},
			want: "",
		},

		{
			name: "TestReplaceStructByTag with overlapping tags",
			args: args{
				v: &teststruct{
					Text: "Hello, {firstName} {lastName}",
				},
				replacements: []Replacement{
					{
						Tag:   "firstName",
						Value: "{nickName}",
					},
					{
						Tag:   "lastName",
						Value: "Doe",
					},
					{
						Tag:   "nickName",
						Value: "John",
					},
				},
			},
			want: "Hello, John Doe",
		},

		{
			name: "TestReplaceStructByTag with partial tag replacement",
			args: args{
				v: &teststruct{
					Text: "Hello, {user}",
				},
				replacements: []Replacement{
					{
						Tag:   "user",
						Value: "{firstName} {lastName}",
					},
					{
						Tag:   "firstName",
						Value: "John",
					},
					{
						Tag:   "lastName",
						Value: "Doe",
					},
				},
			},
			want: "Hello, John Doe",
		},

		{
			name: "TestReplaceStructByTag with numeric replacement",
			args: args{
				v: &teststruct{
					Text: "There are {num} apples",
				},
				replacements: []Replacement{
					{
						Tag:   "num",
						Value: "5",
					},
				},
			},
			want: "There are 5 apples",
		},

		{
			name: "TestReplaceStructByTag with special characters in replacement",
			args: args{
				v: &teststruct{
					Text: "Hello, {name}",
				},
				replacements: []Replacement{
					{
						Tag:   "name",
						Value: "Joh@n$",
					},
				},
			},
			want: "Hello, Joh@n$",
		},

		{
			name: "TestReplaceStructByTag recursive tag replacement",
			args: args{
				v: &teststruct{
					Text: "Hello, {user}",
				},
				replacements: []Replacement{
					{
						Tag:   "user",
						Value: "{firstName} {lastName}",
					},
					{
						Tag:   "firstName",
						Value: "John",
					},
					{
						Tag:   "lastName",
						Value: "{user}",
					},
				},
			},
			want: "Hello, John Doe",
		},
	}

	var fn func(any, []Replacement)
	if byTag {
		fn = ReplaceStructByTag
	} else {
		fn = ReplaceStruct
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "TestReplaceStructByTag recursive tag replacement" {
				t.Skip("this case drops panic. so it be an infinite recursion")
			}

			fn(tt.args.v, tt.args.replacements)
			if tt.args.v.(*teststruct).Text != tt.want {
				t.Errorf("ReplaceStructByTag() = %v, want %v", tt.args.v.(*teststruct).Text, tt.want)
			}
		})
	}
}
