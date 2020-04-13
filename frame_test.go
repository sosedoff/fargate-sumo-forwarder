package main

import (
	"bytes"
	"testing"
)

type example struct {
	input    string
	expected []frame
}

func Test_parseFrames(t *testing.T) {
	examples := []example{
		example{
			``,
			[]frame{},
		},
		example{
			`{}`,
			[]frame{
				frame{0, 2},
			},
		},
		example{
			`{"foo":"bar"}{"awesome":"sauce"}`,
			[]frame{
				frame{0, 13},
				frame{13, 32},
			},
		},
	}

	for _, ex := range examples {
		result := parseFrames([]byte(ex.input))

		if len(result) != len(ex.expected) {
			t.Errorf("Expected %v but got %v\n", len(ex.expected), len(result))
		}

		for i, res := range result {
			exres := ex.expected[i]
			if res[0] != exres[0] || res[1] != exres[1] {
				t.Errorf("Expected %v but got %v\n", exres, res)
			}
		}
	}
}

func Test_parseEntries(t *testing.T) {
	input := `{"foo":"bar"}{"awesome":"sauce"}`
	result := parseEntries([]byte(input))

	expected := []entry{
		[]byte(`{"foo":"bar"}`),
		[]byte(`{"awesome":"sauce"}`),
	}

	for idx, e := range result {
		if !bytes.Equal(e, expected[idx]) {
			t.Errorf("Expected [%s] but got [%s]\n", expected[idx], e)
		}
	}
}
