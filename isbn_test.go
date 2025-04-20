package simpleISBN

import (
	"testing"
)

func TestISBN10CheckDigit(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"030640615", "2"}, // from 0306406152
		{"047195869", "7"}, // Wikipedia example
		{"123456789", "X"}, // yields X
	}

	for _, tt := range tests {
		got := isbn10CheckDigit(tt.input)
		if got != tt.want {
			t.Errorf("isbn10CheckDigit(%q) = %q; want %q", tt.input, got, tt.want)
		}
	}
}

func TestISBN13CheckDigit(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"978030640615", "7"}, // from 9780306406157
		{"978186197271", "2"}, // Wikipedia example
	}

	for _, tt := range tests {
		got, err := isbn13CheckDigit(tt.input)
		if err != nil {
			t.Errorf("isbn13CheckDigit(%q) unexpected error: %v", tt.input, err)
			continue
		}
		if got != tt.want {
			t.Errorf("isbn13CheckDigit(%q) = %q; want %q", tt.input, got, tt.want)
		}
	}
}

func TestConvertISBN(t *testing.T) {
	tests := []struct {
		raw     string
		target  ISBNType
		want    string
		wantErr bool
	}{
		// ISBN10 -> ISBN13
		{"0306406152", ISBN13, "9780306406157", false},
		// ISBN13 -> ISBN10
		{"9780306406157", ISBN10, "0306406152", false},
		// invalid checksum converting to ISBN-13
		{"1234567890", ISBN13, "", true},
		// wrong prefix converting to ISBN-10
		{"9791234567897", ISBN10, "", true},
	}

	for _, tt := range tests {
		got, err := convertISBN(tt.raw, tt.target)
		if (err != nil) != tt.wantErr {
			t.Errorf(
				"convertISBN(%q, %v) error = %v; wantErr %v",
				tt.raw,
				tt.target,
				err,
				tt.wantErr,
			)
			continue
		}
		if got != tt.want {
			t.Errorf("convertISBN(%q, %v) = %q; want %q", tt.raw, tt.target, got, tt.want)
		}
	}
}

func TestNewISBN(t *testing.T) {
	tests := []struct {
		input    string
		wantType ISBNType
		want10   string
		want13   string
		wantErr  bool
	}{
		{"0306406152", ISBN10, "0306406152", "9780306406157", false},
		{"9780306406157", ISBN13, "0306406152", "9780306406157", false},
		{"invalid", "", "", "", true},
		{"123456789012", "", "", "", true},
	}

	for _, tt := range tests {
		isbn, err := NewISBN(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("NewISBN(%q) error = %v; wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			continue
		}
		if isbn.InitialType != tt.wantType {
			t.Errorf(
				"NewISBN(%q).InitialType = %v; want %v",
				tt.input,
				isbn.InitialType,
				tt.wantType,
			)
		}
		if isbn.ISBN10Number != tt.want10 {
			t.Errorf(
				"NewISBN(%q).ISBN10Number = %q; want %q",
				tt.input,
				isbn.ISBN10Number,
				tt.want10,
			)
		}
		if isbn.ISBN13Number != tt.want13 {
			t.Errorf(
				"NewISBN(%q).ISBN13Number = %q; want %q",
				tt.input,
				isbn.ISBN13Number,
				tt.want13,
			)
		}
	}
}
