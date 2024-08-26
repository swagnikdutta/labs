package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	expected := "Swagnik"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}

func TestWalkV2(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Swagnik"},
			[]string{"Swagnik"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with int field",
			struct {
				Name string
				age  int
			}{"Swagnik", 31},
			[]string{"Swagnik"},
		},
		{
			"struct with nested fields",
			Person{
				"Swagnik",
				Profile{31, "Kolkata"},
			},
			[]string{"Swagnik", "Kolkata"},
		},
		{
			"pointer to things",
			&Person{
				"Swagnik",
				Profile{31, "Kolkata"},
			},
			[]string{"Swagnik", "Kolkata"},
		},
		{
			"slices inside struct",
			struct {
				Profiles []Profile
			}{
				[]Profile{
					{31, "Kolkata"},
					{33, "Bangalore"},
				},
			},
			[]string{"Kolkata", "Bangalore"},
		},
		{
			"slices straight away",
			[]Profile{
				{31, "Kolkata"},
				{33, "Bangalore"},
			},
			[]string{"Kolkata", "Bangalore"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
