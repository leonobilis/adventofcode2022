package main

import (
	"testing"
)

var testData1 = []struct {
	dec   int
	snafu string
}{
	{1, "1"},
	{2, "2"},
	{3, "1="},
	{4, "1-"},
	{5, "10"},
	{6, "11"},
	{7, "12"},
	{8, "2="},
	{9, "2-"},
	{10, "20"},
	{15, "1=0"},
	{20, "1-0"},
	{2022, "1=11-2"},
	{12345, "1-0---0"},
	{314159265, "1121-1110-1=0"},
}

var testData2 = []struct {
	snafu string
	dec   int
}{
	{"1=-0-2", 1747},
	{"12111", 906},
	{"2=0=", 198},
	{"21", 11},
	{"2=01", 201},
	{"111", 31},
	{"20012", 1257},
	{"112", 32},
	{"1=-1=", 353},
	{"1-12", 107},
	{"12", 7},
	{"1=", 3},
	{"122", 37},
}

func TestDecode(t *testing.T) {
	for _, td := range testData1 {
		check := decode(td.snafu)
		if check != td.dec {
			t.Fatalf("%v, Expected %v got %v", td.dec, td.snafu, check)
		}
	}

	for _, td := range testData2 {
		check := decode(td.snafu)
		if check != td.dec {
			t.Fatalf("%v, Expected %v got %v", td.dec, td.snafu, check)
		}
	}
}

func TestEncode(t *testing.T) {
	for _, td := range testData1 {
		check := encode(td.dec)
		if check != td.snafu {
			t.Fatalf("%v, Expected %v got %v", td.dec, td.snafu, check)
		}
	}

	for _, td := range testData2 {
		check := encode(td.dec)
		if check != td.snafu {
			t.Fatalf("%v, Expected %v got %v", td.dec, td.snafu, check)
		}
	}
}
