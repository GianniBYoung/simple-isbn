package simpleISBN

import (
	"fmt"
	"strconv"
	"strings"
)

type ISBNType string

const (
	ISBN10 ISBNType = "ISBN-10"
	ISBN13 ISBNType = "ISBN-13"
)

// Representation of an ISBN Number (not including ISBN- prefix)
type ISBN struct {
	// NumberWithHyphens   string // not doing for now because hyphenation varies based on publisher/country
	ISBN10Number string
	ISBN13Number string
	InitialType  ISBNType
	Raw          string
}

// Takes an isbn10|13 and returns an `ISBN` struct
func NewISBN(input string) (*ISBN, error) {
	// normalize input: trim space, to lower, remove hyphens and "isbn" prefix
	input = strings.TrimSpace(strings.ToLower(strings.ReplaceAll(input, "-", "")))
	raw := strings.TrimPrefix(input, "isbn")
	raw = strings.ToUpper(raw) // Upercase 'X' required for ISBN10 and ISBN13 only has digits

	isbn := &ISBN{Raw: raw}

	switch len(raw) {
	case 10:
		isbn.InitialType = ISBN10
		isbn.ISBN10Number = raw

		alt, err := ConvertISBN(raw, ISBN13)
		if err != nil {
			return nil, fmt.Errorf("converting %q → ISBN-13: %w", raw, err)
		}
		isbn.ISBN13Number = alt

	case 13:
		isbn.InitialType = ISBN13
		isbn.ISBN13Number = raw

		alt, err := ConvertISBN(raw, ISBN10)
		if err != nil {
			return nil, fmt.Errorf("converting %q → ISBN-10: %w", raw, err)
		}
		isbn.ISBN10Number = alt

	default:
		return nil, fmt.Errorf("invalid ISBN length %d; must be 10 or 13", len(raw))
	}

	return isbn, nil
}

// Converts an isbn number with out the `isbn-` prefix to ISBNtype `t`
func ConvertISBN(raw string, toType ISBNType) (string, error) {

	switch toType {
	case ISBN10:
		// want ISBN‑10 → so input must be a 13‑digit ISBN‑13
		if len(raw) != 13 {
			return "", fmt.Errorf("cannot convert to ISBN-10: input length %d, want 13", len(raw))
		}
		if !strings.HasPrefix(raw, "978") {
			return "", fmt.Errorf("cannot convert ISBN-13 %q → ISBN-10: prefix must be 978", raw)
		}
		core := raw[3:12]
		cd := isbn10CheckDigit(core)
		return core + cd, nil

	case ISBN13:
		// want ISBN‑13 → so input must be a 10‑digit ISBN‑10
		if len(raw) != 10 {
			return "", fmt.Errorf("cannot convert to ISBN-13: input length %d, want 10", len(raw))
		}
		// validate original ISBN-10 checksum
		expected := isbn10CheckDigit(raw[:9])
		actual := string(raw[9])
		if expected != actual {
			return "", fmt.Errorf("invalid ISBN-10 checksum: expected %q, got %q", expected, actual)
		}
		body := "978" + raw[:9]
		cd, err := isbn13CheckDigit(body)
		if err != nil {
			return "", fmt.Errorf("computing ISBN-13 check digit: %w", err)
		}
		return body + cd, nil

	default:
		return "", fmt.Errorf("unknown target ISBNType %q", toType)
	}
}

// isbn13CheckDigit computes the single-digit checksum for a 12‑digit string.
// It returns that last digit as a string.
func isbn13CheckDigit(twelve string) (string, error) {
	if len(twelve) != 12 {
		return "", fmt.Errorf("must supply 12 digits, got %d", len(twelve))
	}
	sum := 0
	for i, r := range twelve {
		d, err := strconv.Atoi(string(r))
		if err != nil {
			return "", fmt.Errorf("non‑digit %q at pos %d", r, i)
		}
		// weights: 1,3,1,3,...
		if i%2 == 0 {
			sum += d
		} else {
			sum += 3 * d
		}
	}
	check := (10 - (sum % 10)) % 10
	return strconv.Itoa(check), nil
}

// isbn10CheckDigit computes the final checksum character for a 9‑digit string.
// It returns "0"–"9" or "X" if the value is 10.
func isbn10CheckDigit(nine string) string {
	sum := 0
	for i, r := range nine {
		d := int(r - '0') // assume valid digit
		sum += (i + 1) * d
	}
	rem := sum % 11
	if rem == 10 {
		return "X"
	}
	return strconv.Itoa(rem)
}

func (isbn ISBN) String() string {
	return fmt.Sprintf(
		"Inital Type: %s\nInital Raw Input: %s\nISBN-10: %s\nISBN-13: %s\n",
		isbn.InitialType,
		isbn.Raw,
		isbn.ISBN10Number,
		isbn.ISBN13Number,
	)
}
