package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type TestCase struct {
	Name     string
	Input    interface{}
	Expected []string
}

func TestWalk(t *testing.T) {
	cases := []TestCase{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{"Alan"},
			Expected: []string{"Alan"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Alan", "London"},
			Expected: []string{"Alan", "London"},
		},
		{
			Name: "Struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Alan", 26},
			Expected: []string{"Alan"},
		},
		{
			Name: "Struct with nested fields",
			Input: struct {
				Name    string
				Profile Profile
			}{"Alan", Profile{26, "London"}},
			Expected: []string{"Alan", "London"},
		},
		{
			Name: "Pointers",
			Input: &struct {
				Name    string
				Profile Profile
			}{"Alan", Profile{26, "London"}},
			Expected: []string{"Alan", "London"},
		},
		{
			Name: "Slices",
			Input: []Profile{
				{26, "London"},
				{27, "Paris"},
				{28, "New York"},
				{29, "Tokyo"},
			},
			Expected: []string{"London", "Paris", "New York", "Tokyo"},
		},
		{
			Name: "Arrays",
			Input: [4]Profile{
				{26, "London"},
				{27, "Paris"},
				{28, "New York"},
				{29, "Tokyo"},
			},
			Expected: []string{"London", "Paris", "New York", "Tokyo"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			testWalkWithCase(t, test)
		})
	}

	t.Run("Maps", func(t *testing.T) {
		input := map[string]Profile{
			"first":  {26, "London"},
			"second": {27, "Paris"},
			"third":  {28, "New York"},
		}

		expected := []string{"London", "Paris", "New York"}
		got := []string{}
		Walk(input, func(x string) { got = append(got, x) })

		for _, item := range expected {
			if !contains(got, item) {
				t.Errorf("wanted %q but missing", item)
			}
		}
	})

	t.Run("Channels", func(t *testing.T) {
		channel := make(chan Profile)

		go func() {
			channel <- Profile{26, "London"}
			channel <- Profile{27, "Paris"}
			channel <- Profile{28, "New York"}
			close(channel)
		}()

		testWalkWithCase(t, TestCase{
			Name:     "Channels",
			Input:    channel,
			Expected: []string{"London", "Paris", "New York"},
		})
	})

	t.Run("Functions", func(t *testing.T) {
		fn := func() (Profile, Profile, Profile) {
			return Profile{26, "London"}, Profile{27, "Paris"}, Profile{28, "New York"}
		}

		testWalkWithCase(t, TestCase{
			Name:     "Functions",
			Input:    fn,
			Expected: []string{"London", "Paris", "New York"},
		})
	})
}

func testWalkWithCase(t *testing.T, test TestCase) {
	t.Helper()
	got := []string{}
	Walk(test.Input, func(x string) { got = append(got, x) })

	if !reflect.DeepEqual(got, test.Expected) {
		t.Errorf("got %v, want %v", got, test.Expected)
	}
}

func contains(s []string, x string) bool {
	for _, item := range s {
		if item == x {
			return true
		}
	}
	return false
}
